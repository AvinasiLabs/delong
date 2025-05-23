package db

import (
	"context"
	"delong/pkg"
	"delong/pkg/aesgcm"
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"

	ma "github.com/multiformats/go-multiaddr"
)

type IpfsStore struct {
	ipfsApi *rpc.HttpApi
}

func NewIpfsStore(ipfsApiAddr string) (*IpfsStore, error) {
	addr, err := ma.NewMultiaddr(ipfsApiAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse IPFS API address: %v", err)
	}

	ipfsApi, err := rpc.NewApi(addr)
	if err != nil {
		return nil, fmt.Errorf("failed to new ipfs api: %v", err)
	}

	return &IpfsStore{
		ipfsApi: ipfsApi,
	}, nil
}

func (i *IpfsStore) Upload(ctx context.Context, fd []byte) (string, error) {
	f := files.NewBytesFile(fd)
	p, err := i.ipfsApi.Unixfs().Add(ctx, f)
	if err != nil {
		return "", err
	}
	return p.RootCid().String(), nil
}

// UploadStream uploads an io.Reader stream to IPFS without buffering it all in memory.
func (i *IpfsStore) UploadStream(ctx context.Context, r io.Reader) (string, error) {
	f := files.NewReaderFile(r)
	node, err := i.ipfsApi.Unixfs().Add(ctx, f)
	if err != nil {
		return "", err
	}
	return node.RootCid().String(), nil
}

func (i *IpfsStore) UploadEncrypted(ctx context.Context, rawFile []byte, key []byte) (string, error) {
	combined, err := pkg.EncryptGCM(rawFile, key)
	if err != nil {
		return "", err
	}

	cid, err := i.Upload(ctx, combined)
	if err != nil {
		return "", err
	}

	return cid, nil
}

// UploadEncryptedStream reads plaintext from r, encrypts it in fixed-size chunks,
// and uploads the resulting ciphertext stream to IPFS. It returns the root CID.
func (i *IpfsStore) UploadEncryptedStream(ctx context.Context, r io.Reader, key []byte, chunkSize int) (string, error) {
	// Create a pipe, write encrypted data into pw,
	// and read it from pr to feed into IPFS.
	pr, pw := io.Pipe()

	ew, err := aesgcm.NewWriter(pw, key, chunkSize)
	if err != nil {
		pw.Close()
		return "", fmt.Errorf("failed to create encrypt writer: %w", err)
	}

	go func() {
		defer pw.Close()
		// Copy plaintext from r into the encryptor.
		if _, err := io.Copy(ew, r); err != nil {
			// Propagate error to reader side.
			pw.CloseWithError(fmt.Errorf("encryption error: %w", err))
			return
		}
		ew.Close()
	}()

	f := files.NewReaderFile(pr)
	node, err := i.ipfsApi.Unixfs().Add(ctx, f)
	if err != nil {
		return "", fmt.Errorf("IPFS add failed: %w", err)
	}
	return node.RootCid().String(), nil
}

// Download reads the file with the given CID from IPFS.
func (i *IpfsStore) Download(ctx context.Context, cidStr string) ([]byte, error) {
	r, err := i.DownloadStream(ctx, cidStr)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)

	// data, err := io.ReadAll(r)
	// if err != nil {
	// 	return nil, err
	// }
	// // 调试：把下载下来的原始二进制 dump 到 /tmp/ipfs-<cid>.tar.gz
	// home, err := os.UserHomeDir()
	// if err != nil {
	// 	return nil, err
	// }

	// debugFile := filepath.Join(home, fmt.Sprintf("ipfs-%s.tar.gz", cidStr))
	// if err := os.WriteFile(debugFile, data, 0644); err != nil {
	// 	log.Printf("warning: failed to dump IPFS data: %v", err)
	// } else {
	// 	log.Printf("wrote raw IPFS archive to %s", debugFile)
	// }

	// return data, nil
}

// DownloadStream reads the file with the given CID from IPFS.
func (i *IpfsStore) DownloadStream(ctx context.Context, cidStr string) (io.ReadCloser, error) {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return nil, err
	}

	p := path.FromCid(c)

	node, err := i.ipfsApi.Unixfs().Get(ctx, p)
	if err != nil {
		return nil, err
	}

	f, ok := node.(files.File)
	if !ok {
		return nil, fmt.Errorf("node is not a file")
	}

	return f, nil
}

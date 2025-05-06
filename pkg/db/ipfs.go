package db

import (
	"context"
	"delong/pkg"
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

func (i *IpfsStore) Download(ctx context.Context, cidStr string) ([]byte, error) {
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

	return io.ReadAll(f)
}

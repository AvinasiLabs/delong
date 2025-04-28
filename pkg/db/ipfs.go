package db

import (
	"context"
	"fmt"
	"io"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/kubo/client/rpc"
	"github.com/ipfs/kubo/core/coreiface/options"

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

func (i *IpfsStore) Add(ctx context.Context, fd []byte, opts ...options.UnixfsAddOption) (string, error) {
	f := files.NewBytesFile(fd)
	p, err := i.ipfsApi.Unixfs().Add(ctx, f, opts...)
	if err != nil {
		return "", err
	}
	return p.RootCid().String(), nil
}

func (i *IpfsStore) Get(ctx context.Context, cidStr string) ([]byte, error) {
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

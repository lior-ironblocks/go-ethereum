package __ironblocks__

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
)

type client struct {
	cfg *Config
	c   *rpc.Client
}

func newClient() *client {
	return &client{}
}

func (c *client) init(cfg *Config) error {
	var err error
	c.cfg = cfg
	c.c, err = rpc.Dial(c.cfg.getRPCURL())
	if err != nil {
		return err
	}
	return nil
}

func (c *client) do(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.c.CallContext(ctx, result, method, args...)
}

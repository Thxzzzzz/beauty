package grpcclient

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Option func(c *Client)

func WithAddr(addr string) Option {
	return func(c *Client) {
		c.Addr = addr
	}
}

func WithBlock() Option {
	return func(c *Client) {
		c.DialOpts = append(c.DialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

func WithInsecure() Option {
	return func(c *Client) {
		c.DialOpts = append(c.DialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
}

func WithDialOpts(opts ...grpc.DialOption) Option {
	return func(c *Client) {
		c.DialOpts = append(c.DialOpts, opts...)
	}
}

type Client struct {
	*grpc.ClientConn
	Addr              string
	DialOpts          []grpc.DialOption
	unaryInterceptors []grpc.UnaryClientInterceptor
}

func (c *Client) Close() {
	c.ClientConn.Close()
}

func New(opts ...Option) (*Client, error) {
	c := &Client{
		DialOpts: []grpc.DialOption{
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                time.Second * 20,
				Timeout:             time.Second * 10,
				PermitWithoutStream: true,
			}),
			grpc.WithIdleTimeout(time.Second * 10),
			// grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithChainUnaryInterceptor(),
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	conn, err := grpc.DialContext(context.Background(), c.Addr, c.DialOpts...)
	if err != nil {
		return &Client{}, err
	}
	c.ClientConn = conn
	return c, nil
}

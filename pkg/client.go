//
// Copyright (C) 2024 dszi
//
// This file may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.
// Repository: https://github.com/dszi/go-tron
//

package pkg

import (
	"context"
	"fmt"
	"time"

	"github.com/dszi/go-tron/pb/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// GrpcClient represents a gRPC client for the TRON API.
type GrpcClient struct {
	Address     string
	Conn        *grpc.ClientConn
	Client      api.WalletClient
	grpcTimeout time.Duration
	opts        []grpc.DialOption
	apiKey      string
}

// Option defines a function type for configuring a GrpcClient.
type Option func(*GrpcClient)

// WithTimeout sets the timeout for gRPC requests.
func WithTimeout(timeout time.Duration) Option {
	return func(c *GrpcClient) {
		c.grpcTimeout = timeout
	}
}

// WithAPIKey sets the API key for authenticated requests.
func WithAPIKey(key string) Option {
	return func(c *GrpcClient) {
		c.apiKey = key
	}
}

// WithDialOptions sets additional gRPC dial options.
func WithDialOptions(opts ...grpc.DialOption) Option {
	return func(c *GrpcClient) {
		c.opts = opts
	}
}

// NewGrpcClient creates a new GrpcClient with the specified address and options.
func NewGrpcClient(address string, options ...Option) TronClient {
	client := &GrpcClient{
		Address:     address,
		grpcTimeout: 5 * time.Second, // default timeout
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

// Start initializes the gRPC connection.
func (g *GrpcClient) Start() error {
	if g.Address == "" {
		g.Address = "grpc.trongrid.io:50051"
	}
	conn, err := grpc.Dial(g.Address, g.opts...)
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	g.Conn = conn
	g.Client = api.NewWalletClient(conn)
	return nil
}

// getContext returns a context with timeout and attaches the API key if set.
func (g *GrpcClient) getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), g.grpcTimeout)
	if g.apiKey != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "TRON-PRO-API-KEY", g.apiKey)
	}
	return ctx, cancel
}

// Stop closes the gRPC connection.
func (g *GrpcClient) Stop() {
	if g.Conn != nil {
		_ = g.Conn.Close()
	}
}

// Reconnect stops and restarts the gRPC connection with an optional new URL.
func (g *GrpcClient) Reconnect(url string) error {
	g.Stop()
	if url != "" {
		g.Address = url
	}
	return g.Start()
}

// GetMessageBytes creates a BytesMessage from a byte slice.
func GetMessageBytes(m []byte) *api.BytesMessage {
	return &api.BytesMessage{Value: m}
}

// GetMessageNumber creates a NumberMessage from an int64.
func GetMessageNumber(n int64) *api.NumberMessage {
	return &api.NumberMessage{Num: n}
}

// GetPaginatedMessage creates a PaginatedMessage with offset and limit.
func GetPaginatedMessage(offset, limit int64) *api.PaginatedMessage {
	return &api.PaginatedMessage{
		Offset: offset,
		Limit:  limit,
	}
}

// validateTx checks the transaction response for errors.
func validateTx(tx *api.TransactionExtention) error {
	if proto.Size(tx) == 0 {
		return fmt.Errorf("bad transaction")
	}
	if tx.GetResult().GetCode() != 0 {
		return fmt.Errorf("%s", tx.GetResult().GetMessage())
	}
	return nil
}

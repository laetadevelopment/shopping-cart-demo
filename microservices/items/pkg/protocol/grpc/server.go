package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/laetadevelopment/shopping-cart-demo/microservices/items/pkg/api/v1"
	"github.com/laetadevelopment/shopping-cart-demo/microservices/items/pkg/logger"
	"github.com/laetadevelopment/shopping-cart-demo/microservices/items/pkg/protocol/grpc/middleware"
)

func RunServer(ctx context.Context, v1API v1.ItemServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}

	opts = middleware.AddLogging(logger.Log, opts)

	server := grpc.NewServer(opts...)
	v1.RegisterItemServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}

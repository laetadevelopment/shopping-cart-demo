package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"source.cloud.google.com/laeta-shopping-cart-demo/ShoppingCartDemo/pkg/api/v1"
	"source.cloud.google.com/laeta-shopping-cart-demo/ShoppingCartDemo/pkg/logger"
	"source.cloud.google.com/laeta-shopping-cart-demo/ShoppingCartDemo/pkg/protocol/grpc/middleware"
)

func RunServer(ctx context.Context, v1API v1.CartServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}

	opts = middleware.AddLogging(logger.Log, opts)

	server := grpc.NewServer(opts...)
	v1.RegisterCartServiceServer(server, v1API)

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

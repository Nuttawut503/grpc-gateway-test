package main

import (
	"context"
	"gogateway/server"
	"gogateway/server/customerpb"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	restEndpoint = ":8080"
	grpcEndpoint = ":9090"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := runGRPC(ctx, grpcEndpoint); err != nil {
			log.Fatal(err)
		}
	}()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := customerpb.RegisterCustomerHandlerFromEndpoint(context.Background(), mux, grpcEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	srv := http.Server{
		Addr:    restEndpoint,
		Handler: mux,
	}
	go func() {
		<-ctx.Done()
		stop()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Gateway server is running...")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Println("Gateway server is closed...")
}

func runGRPC(ctx context.Context, port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	srv := grpc.NewServer()
	customerpb.RegisterCustomerServer(srv, &server.Server{})
	go func() {
		<-ctx.Done()
		srv.Stop()
	}()

	log.Println("gRPC server is running...")
	if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		return err
	}
	log.Println("gRPC server is closed...")
	return nil
}

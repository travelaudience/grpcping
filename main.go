package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"time"
)

func main() {
	var (
		address, service, certFile string
		timeout time.Duration
		conn                       *grpc.ClientConn
		clientCreds                credentials.TransportCredentials
		err                        error
	)

	flag.StringVar(&address, "address", "", "Service address.")
	flag.StringVar(&service, "service", "", "Service name.")
	flag.StringVar(&certFile, "certfile", "", "Path to a certificate.")
	flag.DurationVar(&timeout, "timeout", 5 *time.Second, "Maximum amount of time application will wait for response.")
	flag.Parse()

	if certFile == "" {
		conn, err = grpc.Dial(address, grpc.WithInsecure())
	} else {
		clientCreds, err = credentials.NewClientTLSFromFile(certFile, "")
		fail("tls initialization failed", err)

		conn, err = grpc.Dial(address, grpc.WithTransportCredentials(clientCreds))
	}
	fail("connection cannot be established", err)

	client := healthpb.NewHealthClient(conn)

	ctx,cancel:= context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := client.Check(ctx, &healthpb.HealthCheckRequest{
		Service: service,
	})
	fail("health check request failed", err)

	fmt.Println(resp.String())
}

func fail(msg string, err error) {
	if err != nil {
		fmt.Println(msg + ": " + err.Error())
		os.Exit(1)
	}
}

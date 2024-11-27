package servers

import (
	"context"
	"log"
	"net"
	"peerbill-user-api/config"
	"peerbill-user-api/gapi"
	"peerbill-user-api/pb"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

/**
 * StartGrpcServer initializes and starts the gRPC server to serve
 * requests over HTTP/2. It does the following:
 *
 * 1. Creates an instance of the `PeerbillUserServer` implementation.
 * 2. Sets up a logger for request tracking with a gRPC unary interceptor.
 * 3. Registers the server implementation with the gRPC server.
 * 4. Registers the reflection service to allow for introspection.
 * 5. Starts the gRPC server in a separate goroutine to listen for incoming requests.
 * 6. Starts a second goroutine to listen for shutdown signals from the provided context (`ctx`),
 *    gracefully shutting down the server when the context is done.
 */
func StartGrpcServer(group *errgroup.Group, ctx context.Context, config config.Config) {
	// an implementation of the PeerbillUserServer
	server := gapi.NewGServer(config)

	// logger for tracking requests sent to the server
	options := grpc.UnaryInterceptor(gapi.GrpcLogger)

	// create a new grpc server and register with own implementation
	gServer := grpc.NewServer(options)
	pb.RegisterPeerbillUserServer(gServer, server)

	// register reflction service
	reflection.Register(gServer)

	/**
	* start a separate go routine to listen and serve requests
	* on the grpc server
	 */
	listener, err := net.Listen("tcp", config.GRPCServerAddr)
	if err != nil {
		log.Fatal(err)
	}
	group.Go(func() error {
		log.Print("Grpc server running on: ", config.GRPCServerAddr)
		err := gServer.Serve(listener)
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	/**
	* start a separate go routine to listen for shutdown signals
	* by calling the Done method on the ctx
	 */
	group.Go(func() error {
		<-ctx.Done()
		log.Print("Grpc server gracefully shutting down...")

		gServer.GracefulStop()
		log.Print("goodbye")

		return nil
	})
}

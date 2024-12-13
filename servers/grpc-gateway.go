package servers

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/ebukacodes21/peerbill-user-api/config"
	db "github.com/ebukacodes21/peerbill-user-api/db/sqlc"
	_ "github.com/ebukacodes21/peerbill-user-api/doc/statik"
	"github.com/ebukacodes21/peerbill-user-api/gapi"
	"github.com/ebukacodes21/peerbill-user-api/pb"
)

/**
* The Grpc Gateway will accept client requests and re route the grpc server
* The RegisterPeerbillUserHandlerServer is important as it used to register
* the server and a gateway
 */
func StartGrpcGateway(group *errgroup.Group, ctx context.Context, config config.Config, repository db.DatabaseContract) {
	// an implementation of the PeerbillUserServer
	server := gapi.NewGServer(config, repository)

	// marshaler options to pass into the grpc mux handler
	options := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// set up gateway/grpc mux by calling NewServeMux method on the runtime package
	grpcMux := runtime.NewServeMux(options)
	pb.RegisterPeerbillUserHandlerServer(ctx, grpcMux, server)

	/**
	* register a HTTP1 mux that will listen for requests and pass
	* to the grpc server for processing/handling
	 */
	httpMux := http.NewServeMux()
	httpMux.Handle("/", grpcMux)

	// embed static swagger files in the server memory for efficiency
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Serve swagger documentation over HTTP.
	httpMux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(statikFS)))

	// set up cors to allow requests from trusted origins
	c := cors.New(cors.Options{
		AllowedOrigins: config.AllowedOrigins,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"Authorization",
			"Content-Type",
		},
	})
	handler := c.Handler(gapi.HttpLogger(httpMux))

	// instantiate a HTTP2 server for enhanced functionality
	httpServer := &http.Server{
		Handler: handler,
		Addr:    config.HTTPServerAddr,
	}

	/**
	* start a separate go routine to listen and serve requests
	* on the gateway
	 */
	group.Go(func() error {
		log.Print("Gateway server running on: ", config.HTTPServerAddr)
		err = httpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}

		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		log.Print("Gateway server gracefully shutting down...")

		err = httpServer.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Print("goodbye")
		return nil
	})
}

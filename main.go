package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"peerbill-user-api/config"
	"peerbill-user-api/servers"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// shutdown signals
var signals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func main() {
	/**
	* loads environment values from the .env file
	* this method returns a configuration struct
	* and a potential error
	 */
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	/**
	 * NotifyContext of the signal package listens for
	 * OS signals by accepting a context and a list of
	 * signals to watch for. It returns a new context
	 * that will be cancelled when any of the specified
	 * signals are received, and a stop function to
	 * release resources when done.
	 */
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	/**
	* calling the WithContext method on the errorgroup package
	* will return a group and a context associated with the
	* context above. The group will be used to run a multiple
	* async tasks in parallel. a group will terminate when a ctx
	* is cancelled or a task reports an error
	 */
	group, ctx := errgroup.WithContext(ctx)

	/**
	* start the grpc server
	 */
	servers.StartGrpcServer(group, ctx, config)

	// wait for all methods to return before exiting the main func
	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

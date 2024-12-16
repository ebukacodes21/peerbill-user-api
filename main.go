package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebukacodes21/peerbill-user-api/config"
	db "github.com/ebukacodes21/peerbill-user-api/db/sqlc"
	"github.com/ebukacodes21/peerbill-user-api/gapi"
	"github.com/ebukacodes21/peerbill-user-api/servers"

	_ "github.com/lib/pq"
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

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewRepository(conn)
	server := gapi.NewGServer(config, repository)

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
	* start the grpc, gateway servers
	 */
	servers.RunDBMigration(config.MigrationURL, config.DBSource)
	servers.StartGrpcServer(group, ctx, config, repository)
	servers.StartGrpcGateway(group, ctx, config, repository)
	servers.RunWebSocket(group, ctx, config, *server)

	// wait for all methods to return before exiting the main func
	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

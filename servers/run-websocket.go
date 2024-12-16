package servers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	tp "github.com/ebukacodes21/peerbill-trader-api/pb"
	"github.com/ebukacodes21/peerbill-user-api/config"
	"github.com/ebukacodes21/peerbill-user-api/gapi"
	"github.com/ebukacodes21/peerbill-user-api/pb"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GetRatesMessage struct {
	Crypto string `json:"crypto"`
	Fiat   string `json:"fiat"`
}

type GetTradersMessage struct {
	Crypto string `json:"crypto"`
	Fiat   string `json:"fiat"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func RunWebSocket(group *errgroup.Group, ctx context.Context, config config.Config, server gapi.GServer) {
	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/ws/get-rates", func(w http.ResponseWriter, r *http.Request) {
		handleGetRates(w, r, ctx, server)
	})

	httpMux.HandleFunc("/ws/get-traders", func(w http.ResponseWriter, r *http.Request) {
		handleGetTraders(w, r, ctx)
	})

	// Set up your HTTP server
	httpServer := &http.Server{
		Handler: httpMux,
		Addr:    config.WebsocketAddr,
	}

	// Start the WebSocket server in a goroutine
	group.Go(func() error {
		log.Print("WebSocket server listening on ", config.WebsocketAddr)
		err := httpServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return nil
	})

	// Gracefully shutdown WebSocket server
	group.Go(func() error {
		<-ctx.Done()
		log.Print("Gracefully shutting down WebSocket server...")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Print("WebSocket server shutdown completed.")
		return nil
	})
}

// handle get rates
func handleGetRates(w http.ResponseWriter, r *http.Request, ctx context.Context, server gapi.GServer) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// Handle the WebSocket connection
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var message GetRatesMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		args := pb.RateRequest{
			Crypto: message.Crypto,
			Fiat:   message.Fiat,
		}

		result, err := server.GetRates(ctx, &args)
		if err != nil {
			log.Println("Error fetching rates:", err)
			continue
		}

		responseMsg, err := json.Marshal(result)
		if err != nil {
			log.Println("Error marshalling response:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseMsg); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

// handle get traders
func handleGetTraders(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// Handle the WebSocket connection
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		traderServer, err := grpc.NewClient("0.0.0.0:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatal("unable to connect to server")
		}

		var message GetTradersMessage
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}
		client := tp.NewPeerbillTraderClient(traderServer)
		args := &tp.GetTradersRequest{
			Crypto: message.Crypto,
			Fiat:   message.Fiat,
		}

		resp, err := client.GetTraders(ctx, args)
		if err != nil {
			log.Fatal("unable to fetch traders")
		}

		responseMsg, err := json.Marshal(resp)
		if err != nil {
			log.Println("Error marshalling response:", err)
			continue
		}

		if err := conn.WriteMessage(websocket.TextMessage, responseMsg); err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

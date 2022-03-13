package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/locmai/yuta/common/jetstream"
	"github.com/locmai/yuta/components/messaging/clients"
	"github.com/locmai/yuta/components/messaging/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.ParseFlags()
	router := mux.NewRouter()

	jetstream.Prepare(&cfg.JetStream)

	if cfg.Metrics.Enabled {
		router.Path("/metrics").Handler(promhttp.Handler())
	}
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		WriteTimeout: time.Duration(cfg.Server.Timeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Server.Timeout) * time.Second,
	}
	serverStartTime := time.Now().UnixMilli()
	logrus.Printf("Start time recorded %d", serverStartTime)
	logrus.Printf("Metrics enabled: %v", cfg.Metrics.Enabled)

	var nluClient []clients.NluClient
	for _, nluClientConfig := range cfg.NluClients {
		switch clientType := nluClientConfig.Type; clientType {
		case config.DiaglogflowClientType:
			client, err := clients.NewDiaglogflowClient(nluClientConfig)
			if err != nil {
				panic(err)
			}
			nluClient = append(nluClient, client)
		default:
			logrus.Printf("Client is not implemented %s", clientType)
		}
	}

	for _, chatClientConfig := range cfg.ChatClients {
		switch clientType := chatClientConfig.ChatClientType; clientType {
		case config.MatrixType:
			// TODO: Use only one nluclient at the moment
			// So it would be the first one
			client, err := clients.NewMatrixClient(chatClientConfig, serverStartTime, nluClient[0])
			if err != nil {
				panic(err)
			}
			go client.StartSyncer()
		default:
			logrus.Printf("Client is not implemented %s", clientType)
		}

	}

	logrus.Println("Server started")
	logrus.Fatal(srv.ListenAndServe())
}

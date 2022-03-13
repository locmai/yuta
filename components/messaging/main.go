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
	"github.com/locmai/yuta/components/messaging/producers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.ParseFlags()
	router := mux.NewRouter()

	js := jetstream.Prepare(&cfg.JetStream)

	if cfg.CommonConfig.Metrics.Enabled {
		router.Path("/metrics").Handler(promhttp.Handler())
	}
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	coreProducer := &producers.CoreProducer{
		JetStream: js,
		Topic:     cfg.CommonConfig.JetStream.TopicFor(jetstream.ActionableItemEvent),
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", cfg.CommonConfig.Server.Host, cfg.CommonConfig.Server.Port),
		WriteTimeout: time.Duration(cfg.CommonConfig.Server.Timeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.CommonConfig.Server.Timeout) * time.Second,
	}
	serverStartTime := time.Now().UnixMilli()
	logrus.Printf("Start time recorded %d", serverStartTime)
	logrus.Printf("Metrics enabled: %v", cfg.CommonConfig.Metrics.Enabled)

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
			client, err := clients.NewMatrixClient(chatClientConfig, serverStartTime, nluClient[0], *coreProducer)
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

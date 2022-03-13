package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/locmai/yuta/common/jetstream"
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

	logrus.Println("Server started")
	logrus.Fatal(srv.ListenAndServe())
}

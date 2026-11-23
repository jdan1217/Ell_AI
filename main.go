package main

import (
	"log"
	"net/http"

	"github.com/jdan1217/Ell_AI/internal/api"
	"github.com/jdan1217/Ell_AI/internal/config"
	"github.com/jdan1217/Ell_AI/internal/traffic"

	"github.com/jdan1217/Ell_AI/internal/proxy"
)

func main() {
	log.Println("Ell_AI proxy is starting...")
	configs := config.DefaultConfig()

	log.Printf("Proxy listening on %s:%d", configs.Host, configs.Port)
	log.Printf("Interception Enabled: %t", configs.InterceptionEnabled)
	log.Printf("Log type: %s", configs.LogLevel)

	history := &traffic.History{}
	apiServer := api.NewServer(history)

	go func() {
		log.Println("Management API listening on 127.0.0.1:8081")

		if err := http.ListenAndServe("127.0.0.1:8081", apiServer.Handler()); err != nil {
			log.Fatalf("Management API failed: %v", err)
		}
	}()

	if err := proxy.StartServer(configs, history); err != nil {
		log.Fatal(err)
	}

}

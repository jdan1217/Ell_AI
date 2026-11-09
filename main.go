package main

import (
	"log"

	"github.com/jdan1217/Ell_AI/internal/config"

	"github.com/jdan1217/Ell_AI/internal/proxy"
)

func main() {
	log.Println("Ell_AI proxy is starting...")
	configs := config.DefaultConfig()

	log.Printf("Proxy listening on %s:%d", configs.Host, configs.Port)
	log.Printf("Interception Enabled: %t", configs.InterceptionEnabled)
	log.Printf("Log type: %s", configs.LogLevel)

	err := proxy.StartServer(configs)
	if err != nil {
		log.Fatal(err)
	}

}

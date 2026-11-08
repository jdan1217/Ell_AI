package main

import (
	"log"

	"github.com/jdan1217/Ell_AI/internal/config"
)

func main() {
	log.Println("Ell_AI proxy is starting...")
	configs := config.DefaultConfig()
	log.Println(configs)
}

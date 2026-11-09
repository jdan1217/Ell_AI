package proxy

import (
	"fmt"
	"net"

	"github.com/jdan1217/Ell_AI/internal/config"

	"log"
)

func StartServer(cfg *config.Config) error {
	proto := "tcp"
	host := cfg.Host
	port := cfg.Port
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen(proto, address)
	if err != nil {
		return err
	}
	defer listener.Close()
	log.Printf("Listening using %s on address %s", proto, address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		log.Printf("Connection made from %s", conn.RemoteAddr())
		conn.Close()
	}
}

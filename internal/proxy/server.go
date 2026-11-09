package proxy

import (
	"fmt"
	"net"

	"github.com/jdan1217/Ell_AI/internal/config"

	"bufio"
	"log"
	"net/http"
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Printf("Connection made from %s", conn.RemoteAddr())
	log.Printf("HTTP Method: %s \n URL: %s", req.Method, req.Host)
}

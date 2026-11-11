package proxy

import (
	"fmt"
	"net"

	"github.com/jdan1217/Ell_AI/internal/config"

	"bufio"
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
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	req, err := parseRequest(reader)
	if err != nil {
		log.Printf("Failed to parse HTTP request: %v", err)
		return
	}
	log.Printf("Connection made from %s", conn.RemoteAddr())
	log.Printf("HTTP Method: %s \n URL: %s", req.Method, req.URL.String())
	log.Printf("Request URL Scheme: %s", req.URL.Scheme)
	log.Printf("Request URL Host: %s", req.URL.Host)
	log.Printf("Request URL Path: %s", req.URL.Path)
	for name, values := range req.Header {
		for _, value := range values {
			log.Printf("Header : %s %s\n", name, value)
		}
	}
}

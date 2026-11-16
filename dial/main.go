package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "example.com:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Printf("TCP Connection established at %s\n", conn.RemoteAddr())
	request := []byte("GET / HTTP/1.1\r\nHost: example.com\r\nConnection: close\r\n\r\n")
	bytesWritten, err := conn.Write(request)
	if err != nil {
		log.Fatal(err)
	}
	if bytesWritten != len(request) {
		log.Fatal("Incomplete request write")
	}
	fmt.Printf("Sent %d bytes", bytesWritten)
	buffer := make([]byte, 4096)

	for {
		bytesRead, err := conn.Read(buffer)

		if bytesRead > 0 {
			fmt.Print(string(buffer[:bytesRead]))
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}
}

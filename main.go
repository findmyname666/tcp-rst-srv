package main

import (
	"flag"
	"log"
	"net"
)

var (
	addr string
	port string
)

func main() {
	flag.StringVar(&port, "port", "", "Server listen port.")
	flag.StringVar(&addr, "addr", "127.0.0.1", "Server listen address.")
	flag.Parse()

	if port == "" {
		log.Fatalf("argument -port, server listen port, is required")
	}

	log.Printf("Starting TPC server on '%s:%s'", addr, port)

	// Create TCP listener.
	l, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		log.Fatalf("listener failed: %s", err)
	}
	defer l.Close()

	for {
		// Accept new connections.
		c, err := l.Accept()
		if err != nil {
			log.Fatalf("unable to accept a new connection: %s", err)
		}

		log.Printf("TCP connection received from: %s", c.RemoteAddr())

		// Use SetLinger to force close the connection.
		err = c.(*net.TCPConn).SetLinger(0)
		if err != nil {
			log.Printf("failed to reset connection: %s", err)
		}

		c.Close()
	}
}

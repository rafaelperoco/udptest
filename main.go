// Language: go
// Path: main.go
// envs: UDP_PORT
// a UDP server that echoes back any data it receives
// log all to stdout

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

type udpServer struct {
	conn *net.UDPConn
}

func (s *udpServer) listen() {
	defer s.conn.Close()
	buf := make([]byte, 1024)
	for {
		n, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("received %s from %s", string(buf[:n]), addr)
		_, err = s.conn.WriteToUDP(buf[:n], addr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	port := os.Getenv("UDP_PORT")
	if port == "" {
		port = "8080"
	}
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listening on %s", conn.LocalAddr())
	s := udpServer{conn}
	s.listen()
}
// Language: go
// env: UDP_SERVER, UDP_PORT and UDP_MESSAGE
// A UDP client that sends a message to a UDP server and prints the response
// log all to stdout

package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	port := os.Getenv("UDP_PORT")
	if port == "" {
		port = "8080"
	}
	addr, err := net.ResolveUDPAddr("udp", os.Getenv("UDP_SERVER")+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	message := os.Getenv("UDP_MESSAGE")
	if message == "" {
		message = "hello world"
	}
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("received %s from %s", string(buf[:n]), conn.RemoteAddr())
}
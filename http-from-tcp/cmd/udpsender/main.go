package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = ":42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("error creating UDP address: %v\n", err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("error creating udp connection: %v\n", err)
	}
	defer conn.Close()
	fmt.Println("Listening for UDP traffic on", port)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading line: %v\n", err)
		}
		if _, err := conn.Write([]byte(line)); err != nil {
			fmt.Printf("error writing to udp: %v\n", err)
		}
	}
}

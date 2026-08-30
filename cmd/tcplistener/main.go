package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %v\n", err)
	}
	defer listener.Close()
	fmt.Println("Listening for TCP traffic on", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("couldn't accept connection: %v", err)
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		ch := getLinesChannel(conn)
		for line := range ch {
			fmt.Println(line)
		}
		fmt.Println("Connection to ", conn.RemoteAddr(), "closed")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)
		data := make([]byte, 8)
		line := ""
		for {
			n, err := f.Read(data)
			if err != nil {
				if errors.Is(err, io.EOF) {
					if line != "" {
						ch <- line
					}
					break
				}
				fmt.Printf("error: %v\n", err)
			}
			d := string(data[:n])
			lines := strings.Split(d, "\n")
			line += lines[0]
			for _, part := range lines[1:] {
				ch <- line
				line = part
			}
		}
	}()
	return ch
}

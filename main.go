package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <- chan string {
	out := make(chan string, 1)

	go func() {
		defer close(out)
		defer f.Close()
		
		s := ""

		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}
			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				s += string(data[:i])
				data = data[i + 1:]
				out <- s
				s = ""
			}

			s += string(data)

		}

		if len(s) != 0 {
			out <- s
		}
	}()

	return out
}

func main() {

	listener, err := net.Listen("tcp",":42069")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Println("A connection has been accepted")
		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
		fmt.Println("channel is closed")
	}
}
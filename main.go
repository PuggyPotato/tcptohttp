package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

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
			fmt.Printf("read: %s\n", s)
			s = ""
		}

		s += string(data)

	}

	if len(s) != 0 {
		fmt.Printf("read: %s\n", s)
	}
}
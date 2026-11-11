package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("%s\n", line)
	}
}
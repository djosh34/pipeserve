package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var content []byte
var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to serve on")
}

func main() {
	flag.Parse()

	// Check if there's piped input
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Error: No piped input detected")
		fmt.Println("Usage: cat file.txt | pipeserve [port]")
		os.Exit(1)
	}

	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	var err error
	content, err = io.ReadAll(reader)
	if err != nil {
		log.Fatal("Error reading from stdin:", err)
	}

	// Override port if provided as argument
	if flag.NArg() > 0 {
		_, err := fmt.Sscanf(flag.Arg(0), "%d", &port)
		if err != nil {
			log.Fatal("Invalid port number")
		}
	}

	// Handle all paths
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	})

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Serving on http://localhost%s\n", addr)
	fmt.Println("Press Ctrl+C to stop")
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

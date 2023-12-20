package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalf("error: %s", err)
		/*
			log.Fatalf("error: %s", err)
			os.Exit(1)
		*/
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	//fmt.Printf("Header: %s\n", resp.Header.Get("Content-Type"))

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: copy - %s", err)
	}
}

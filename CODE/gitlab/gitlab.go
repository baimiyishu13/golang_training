package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://gitlab.com/users/baimiyishu13")
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
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

	//if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	//	log.Fatalf("error: copy - %s", err)
	//}

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: %s", err)
	}

	//fmt.Println(r)
	fmt.Printf("%#v\n", r)
}

func jsonInfo(url string) (string, int, error) {
	//TODO: you code goes here

	return "", 0, nil
}
1
type Reply struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

/* Json -> Go
true/false	-	true/false
string	-	string
null	-	nil
number	-	int ... uint ... float ...
array	-	[]any	([]interface{})
object	-	map[string]an,struct

JSON - io.Reader->	Go:JSON.Decode
JSON - []byte	-> GO: JSON.Unmarshal

Go ->io.Writer-> JSON: json.Encoder
Go -[]byte -> Go: json.Marshal
*/

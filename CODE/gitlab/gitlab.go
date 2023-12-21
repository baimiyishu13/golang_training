package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	url2 "net/url"
)

func main() {
	title, id, err := jsonInfo("3")
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("jsoninfo: 👀 title: %s,id: %d\n", title, id)

}

// main中所有内容合并到jsonInfo中
func jsonInfo(num string) (string, int, error) {
	//TODO: you code goes here
	url := "https://jsonplaceholder.typicode.com/todos/" + url2.PathEscape(num)
	println("URL: ", url)
	resp, err := http.Get(url)

	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	fmt.Printf("Header: %s\n", resp.Header.Get("Content-Type"))

	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}

	return r.Title, r.Id, nil
}

type Reply struct {
	UserID    int    `json:"UserID"`
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

package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	//var f *os.File
	//defer f.Close()
	fileName := "http.log.gz"
	shasum(fileName)
}

// 数字签名 cat http.log.gz| gunzip | shasum
func shasum(fileName string) (string, error) {
	// todo: check file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	// 减压
	r, err := gzip.NewReader(file)

	//io.CopyN(os.Stdout, r, 100)
	n, err := io.CopyN(os.Stdout, r, 200)
	if err != nil {
		return "", err
	}
	fmt.Println(n)
	return "", err
}

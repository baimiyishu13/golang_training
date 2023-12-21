package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//var f *os.File
	//defer f.Close()
	fileName := "http.log.gz"
	sig, err := shasum(fileName)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
	fmt.Println(sig)
}

// 数字签名 cat http.log.gz| gunzip | shasum
func shasum(fileName string) (string, error) {
	// todo: check file
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	// 减压
	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer file.Close()

	//io.CopyN(os.Stdout, r, 100)
	n, err := io.CopyN(os.Stdout, r, 100)
	if err != nil {
		return "", err
	}
	fmt.Println(n)

	// 校验
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), err
}

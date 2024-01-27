package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	file, err := os.OpenFile("gear", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := io.LimitReader(rand.Reader, 2048)
	if _, err := io.Copy(file, reader); err != nil {
		panic(err)
	}
}

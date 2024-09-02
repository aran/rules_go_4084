package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	c := 32 * 6 / 8
	out := os.Stdout

	encoded := base64.NewEncoder(base64.URLEncoding, out)
	defer encoded.Close()
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "rand error: ", err)
		os.Exit(2)
	}
	_, err = encoded.Write(b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "write error: ", err)
		os.Exit(3)
	}
}

//go:build !js

package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please specify a demo to run")
	}
	run(os.Args[1])
}

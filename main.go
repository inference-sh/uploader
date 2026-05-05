package main

import (
	"log"

	"inference.sh/uploader/internal/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

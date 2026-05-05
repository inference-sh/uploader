package main

import (
	"log"

	"github.com/inference-sh/uploader/internal/cmd"
)

func main() {
	if err := cmd.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/bluesky6529/go_ucloud/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}

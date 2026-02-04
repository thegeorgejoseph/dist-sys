package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/thegeorgejoseph/dist-sys/pkg/mael"
)

func main() {
	server := mael.Server{Node: maelstrom.NewNode()}

	server.Node.Handle("generate", server.GenerateHandler)

	if err := server.Node.Run(); err != nil {
		log.Fatal(err)
	}

}

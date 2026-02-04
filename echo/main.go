package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"github.com/thegeorgejoseph/dist-sys/pkg/mael"
)

func main() {
	s := &mael.Server{Node: maelstrom.NewNode()}

	s.Node.Handle("echo", s.EchoHandler)

	// basically the runner where jepsen will send messages and STDIN
	// and our handler will send messages back as STDOUT
	if err := s.Node.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"encoding/json"
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// server that will be run on every maelstrom node
type server struct {
	n *maelstrom.Node
}

/*
handler for the echo functionality
node.Reply hadnles all the stuff that the jepsen.io needs such as
src, dest, in_reply_to etc.
we just need to pass in the right body
which is echo_ok in this scenario
*/
func (s *server) echoHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	body["type"] = "echo_ok"
	return s.n.Reply(msg, body)
}

func main() {
	s := &server{n: maelstrom.NewNode()}

	s.n.Handle("echo", s.echoHandler)

	// basically the runner where jepsen will send messages and STDIN
	// and our handler will send messages back as STDOUT
	if err := s.n.Run(); err != nil {
		log.Fatal(err)
	}
}

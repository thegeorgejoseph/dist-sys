package mael

import (
	"encoding/json"

	"github.com/google/uuid"
	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type Server struct {
	Node *maelstrom.Node
}

/*
handler for the echo functionality
node.Reply hadnles all the stuff that the jepsen.io needs such as
src, dest, in_reply_to etc.
we just need to pass in the right body
which is echo_ok in this scenario
*/
func (s *Server) EchoHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	body["type"] = "echo_ok"
	return s.Node.Reply(msg, body)
}

func (s *Server) GenerateHandler(msg maelstrom.Message) error {
	id := uuid.New()
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}
	body["type"] = "generate_ok"
	body["id"] = id

	return s.Node.Reply(msg, body)
}

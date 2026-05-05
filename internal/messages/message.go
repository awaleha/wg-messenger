package messages

import (
	"encoding/json"
	"io"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        string    `json:"id"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Type      string    `json:"type"`
	Body      string    `json:"body"`
	Timestamp time.Time `json:"sent_at"`
}

const MessageTypeChat = "chat"

func NewTextMessage(from, to, body string) Message {
	return Message{
		ID:        uuid.New().String(),
		Type:      "chat",
		From:      from,
		To:        to,
		Body:      body,
		Timestamp: time.Now().UTC(),
	}
}
func Encode(w io.Writer, msg Message) error {
	return json.NewEncoder(w).Encode(msg)
}

func Decode(r io.Reader) (Message, error) {
	var msg Message
	return msg, json.NewDecoder(r).Decode(&msg)
}

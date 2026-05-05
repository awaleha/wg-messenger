package transport

import (
	"fmt"
	"log"
	"net"

	"github.com/awaleha/wg-messenger/internal/messages"
)

func Client(socket string, body string) error {
	//Create tcp socket
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		log.Fatalln(err)
	}

	msg := messages.NewTextMessage("client", "server", body)

	fmt.Println("Sending message", msg.Body, "to", msg.To)
	err = messages.Encode(conn, msg)

	conn.Close()
	return err
}

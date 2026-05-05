package transport

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/awaleha/wg-messenger/internal/messages"
)

func handleConnection(conn net.Conn, f *os.File) {
	defer conn.Close()

	//read from the client
	msg, err := messages.Decode(conn)
	if err != nil {
		fmt.Fprintln(f, "Error reading from user: "+err.Error())
	}

	fmt.Fprintln(f, msg.From, "->", msg.To, ":", msg.Body)
}

func Server(socket string) error {
	ln, err := net.Listen("tcp", socket)
	if err != nil {
		return err
	}
	fmt.Println("<Server> listening on socket -> ", socket)
	defer ln.Close()

	//Output file creation
	f, err := os.Create("server_output.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(f, " Starting server log")

	for {

		//Waiting for new connections
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		//Handle the connection
		go handleConnection(conn, f)
	}
}

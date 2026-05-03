package transport

import (
	"fmt"
	"log"
	"net"
	"os"
)

func handleConnection(conn net.Conn, f *os.File) {
	defer conn.Close()
	//read from the client
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Fprintln(f, "Error reading from user: "+err.Error())
	}

	fmt.Fprintln(f, "<Server> Received Message: ", string(buf[:n]), "From <Client>: ", conn.RemoteAddr())

	fmt.Fprintf(conn, "%s", "echo "+string(buf[:n]))
}

func Server() error {
	fmt.Println("<Server> Starting Server...")
	ln, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		return err
	}
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

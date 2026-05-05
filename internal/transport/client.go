package transport

import (
	"fmt"
	"log"
	"net"
)

func Client(socket string, msg string) error {
	//Create tcp socket
	conn, err := net.Dial("tcp", socket)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("<Client> Sending message to ", socket)
	fmt.Fprintf(conn, "%s", msg)

	//Read response
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("<Client> " + string(buf[:n]))

	conn.Close()
	return nil
}

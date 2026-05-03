package transport

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Client() error {
	//Create tcp socket
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		log.Fatalln(err)
	}

	//Scanner for stdnin
	scanner := bufio.NewReader(os.Stdin)
	fmt.Print("<Client> Please enter message: ")
	line, err := scanner.ReadString('\n')

	fmt.Println("<Client> Sending message...")
	fmt.Fprintf(conn, "%s", line)

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

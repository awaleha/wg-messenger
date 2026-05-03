package main

import (
	"fmt"

	"github.com/awaleha/wg-messenger/internal/transport"
)

func main() {

	fmt.Println("wgmsg: WireGuard messenger")

	//Starting server, returning errors
	errCh := make(chan error, 1)

	go func() {
		errCh <- transport.Server()
	}()

	for {
		transport.Client()
	}

}

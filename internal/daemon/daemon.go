package daemon

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/awaleha/wg-messenger/internal/messages"
)

type Daemon struct {
	logger  *slog.Logger
	logFile *os.File
}

func NewDaemon() (*Daemon, error) {
	f, err := os.OpenFile("daemon.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return nil, err
	}

	logger := slog.New(slog.NewTextHandler(f, nil))
	logger.Info("Starting daemon log")

	return &Daemon{
		logger:  logger,
		logFile: f,
	}, nil

}

func Run(d *Daemon, addr string) error {
	defer d.logFile.Close()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Creating listener
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	d.logger.Info("<daemon> listening on socket -> ", "addr", addr)

	errCh := make(chan error, 1)

	// Catching CTRL+C and closing gracefully
	go func() {
		<-ctx.Done()
		d.logger.Info("shutting down daemon...")
		ln.Close() // forces Accept() to unblock
	}()

	go func() {
		errCh <- ListenLoop(ctx, ln, d)
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errCh:
		return err
	}
}

func ListenLoop(ctx context.Context, ln net.Listener, d *Daemon) error {
	for {
		//Waiting for new connections
		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return nil // expected shutdown
			default:
				return err

			}
		}

		//Handle the connection
		go handleConnection(conn, d)
	}
}

func handleConnection(conn net.Conn, d *Daemon) {
	defer conn.Close()
	//read from the client
	msg, err := messages.Decode(conn)
	if err != nil {
		fmt.Fprintln(d.logFile, "Error reading from user: "+err.Error())
		return
	}

	fmt.Fprintln(d.logFile, "[message]", msg.From, "->", msg.To, ":", msg.Body)
}

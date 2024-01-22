package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func main() {
	flag.DurationVar(&timeout, "timeout", time.Second*10, "возможная задержка подключения")
	flag.Parse()
	client := NewTelnetClient(net.JoinHostPort(flag.Arg(0), flag.Arg(1)), timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func() {
		for {
			err := client.Send()
			if errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stderr, "...EOF")
				os.Exit(0)
			}
			if errors.Is(err, syscall.EPIPE) {
				fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
				os.Exit(0)
			}
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}()
	go func() {
		for {
			err := client.Receive()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}()
	sigCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	sigChannel := sigCtx.Done()
	<-sigChannel
	stop()
	client.Close()
	os.Exit(1)

	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
}

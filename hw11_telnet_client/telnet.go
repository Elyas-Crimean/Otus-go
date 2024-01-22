package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}
type tcImpl struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &tcImpl{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (tc *tcImpl) Connect() error {
	var err error
	tc.conn, err = net.DialTimeout("tcp", tc.address, tc.timeout)
	return err
}

func (tc *tcImpl) Close() error {
	return tc.conn.Close()
}

func (tc *tcImpl) Receive() error {
	_, err := io.Copy(tc.out, tc.conn)
	return err
}

func (tc *tcImpl) Send() error {
	_, err := io.Copy(tc.conn, tc.in)
	return err
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.

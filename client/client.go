package client

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"word-of-wisdom/pow"
)

type Client struct {
	addr string
}

func New(addr string) *Client {
	return &Client{addr}
}

func (c *Client) GetWisdom() (_ string, err error) {
	conn, err := net.Dial("tcp", c.addr)
	if err != nil {
		return
	}
	defer conn.Close()

	_, err = fmt.Fprintf(conn, "HELLO\n")
	if err != nil {
		return
	}

	tokenString, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}

	token, _ := hex.DecodeString(tokenString)

	solution := pow.Solve(token)

	_, err = fmt.Fprintf(conn, "%x\n", solution)
	if err != nil {
		return
	}

	wisdom, err := bufio.NewReader(conn).ReadString('\r')
	if err != nil {
		return
	}
	return wisdom, err
}

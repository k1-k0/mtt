package tools

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

type TCPConnection struct {
	Host string
	Port string
	conn net.Conn
}

func (c *TCPConnection) Connect() error {
	var err error
	c.conn, err = net.Dial("tcp", c.buildAddress())

	if err != nil {
		return errors.New("failure connect")
	}

	return nil
}

func (c *TCPConnection) Send(data []byte) error {
	_, err := c.conn.Write(data)

	if err != nil {
		return errors.New("failure write data")
	}

	return nil
}

func (c *TCPConnection) Get() ([]byte, error) {
	reader := bufio.NewReader(c.conn)
	bytes, err := reader.ReadBytes('\n')

	if err != nil {
		return nil, errors.New("failure read data")
	}

	return bytes, nil
}

func (c *TCPConnection) Close() (err error) {
	err = c.conn.Close()
	return
}

func (c *TCPConnection) buildAddress() string {
	strBuilder := strings.Builder{}
	fmt.Fprintf(&strBuilder, "%s:%s", c.Host, c.Port)
	return strBuilder.String()
}

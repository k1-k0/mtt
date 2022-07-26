package checker

import (
	"fmt"
	"io"
	"mtt/structures"
	"mtt/tools"
	"os"
)

type Checker struct {
	Connection tools.TCPConnection
}

func (c *Checker) Init(conn tools.TCPConnection) {
	c.Connection = conn
}

func (c *Checker) CheckToken(token, scope string) error {
	request := new(structures.Request)

	err := c.Connection.Connect()
	if err != nil {
		return err
	}

	bytes, err := request.BuildRequest(token, scope)
	if err != nil {
		return err
	}

	err = c.Connection.Send(bytes)
	if err != nil {
		return err
	}

	bytes, err = c.Connection.Get()
	if err != nil {
		return err
	}

	response := new(structures.Response)
	err = response.Decode(bytes)
	if err != nil {
		return err
	}

	c.printResult(os.Stdout, response)

	return nil
}

func (c *Checker) printResult(writer io.Writer, response *structures.Response) {
	fmt.Fprintln(writer, response.String())
}

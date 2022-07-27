package checker

import (
	"mtt/structures"
	"mtt/tools"
)

type Checker struct {
	Connection tools.Connection
}

func (c *Checker) Init(conn tools.Connection) {
	c.Connection = conn
}

func (c *Checker) CheckToken(token, scope string) (string, error) {
	request := new(structures.Request)

	err := c.Connection.Connect()
	if err != nil {
		return "", err
	}

	bytes, err := request.BuildRequest(token, scope)
	if err != nil {
		return "", err
	}

	err = c.Connection.Send(bytes)
	if err != nil {
		return "", err
	}

	bytes, err = c.Connection.Get()
	if err != nil {
		return "", err
	}

	response := new(structures.Response)
	err = response.Decode(bytes)
	if err != nil {
		return "", err
	}

	return response.String(), nil
}

package tools

import (
	"errors"
)

type ArgParser struct {
	Host, Port, Token, Scope string
}

func (a *ArgParser) Parse(args []string) error {
	if len(args) != 5 {
		return errors.New("invalid arguments count. expected <host> <port> <token> <scope>")
	}

	a.Host = args[1]
	a.Port = args[2]
	a.Token = args[3]
	a.Scope = args[4]

	return nil
}

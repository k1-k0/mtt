package main

import (
	"fmt"
	"mtt/checker"
	"mtt/tools"
	"os"
)

func main() {
	argParser := new(tools.ArgParser)

	err := argParser.Parse(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	connection := &tools.TCPConnection{Host: argParser.Host, Port: argParser.Port}
	checker := checker.Checker{}
	checker.Init(connection)

	resp, err := checker.CheckToken(argParser.Token, argParser.Scope)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	fmt.Println(resp)
}

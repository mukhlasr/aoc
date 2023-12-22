package main

import (
	"errors"
	"strings"
)

const (
	CmdPrefix = "$ "
)

func isCommand(text string) bool {
	return strings.HasPrefix(text, CmdPrefix)
}

type Command struct {
	Name   string
	Params []string
}

func ParseCommand(text string) (Command, error) {
	if !isCommand(text) {
		return Command{}, errors.New(text + " is not a valid command string")
	}
	cmd := text[len(CmdPrefix):]
	argv := strings.Split(cmd, " ")
	return Command{
		Name:   argv[0],
		Params: argv[1:],
	}, nil
}

package commands

import (
	"os"
	"strings"
)

func ExecuteCommand() (string, error) {
	if len(os.Args) == 1 {
		return "", &MissingCommandError{}
	}
	if len(os.Args) == 2 {
		return "", &MissingSubCommandError{}
	}
	command := os.Args[1]
	subCommand := os.Args[2]
	params := os.Args[3:]
	paramsMap := make(map[string]string)

	for i, v := range os.Args[3:] {
		if i%2 == 0 && len(params) > i+1 {
			paramsMap[strings.Replace(v, "-", "", -1)] = params[i+1]
		}
	}

	switch command {
	case "api":
		switch subCommand {
		case "get-market-orders":
			return ExecuteGetMarketOrdersCommand(paramsMap)
		default:
			return "", &InvalidSubCommandError{Command: command, SubCommand: subCommand}
		}

	default:
		return "", &InvalidCommandError{Command: command}
	}

}

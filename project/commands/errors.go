package commands

import "fmt"

const HowToUseTxt = `
usage: eve-online-tools <command> <subcommand> [parameters]

example: eve-online-tools api get-market-orders --sectorID 10000020 --format csv

To see help text, you can run:

  eve-online-tools help
  eve-online-tools <command> help`

type MissingCommandError struct{}

func (e *MissingCommandError) Error() string {
	return fmt.Sprintf(`missing command`) + "\r\n" + HowToUseTxt
}

type InvalidCommandError struct {
	Command string
}

func (e *InvalidCommandError) Error() string {
	return fmt.Sprintf(`invalid command: "%s"`, e.Command) + "\r\n" + HowToUseTxt
}

type MissingSubCommandError struct{}

func (e *MissingSubCommandError) Error() string {
	return fmt.Sprintf(`missing subcommand`) + "\r\n" + HowToUseTxt
}

type InvalidSubCommandError struct {
	Command    string
	SubCommand string
}

func (e *InvalidSubCommandError) Error() string {
	return fmt.Sprintf(`valid command "%s", but invalid subcommand: "%s"`, e.Command, e.SubCommand) + "\r\n" + HowToUseTxt
}

type UnknownParameterError struct {
	Parameter string
}

func (e *UnknownParameterError) Error() string {
	return fmt.Sprintf(`unknown parameter "%s"`, e.Parameter) + "\r\n" + HowToUseTxt
}

type InvalidParameterValueError struct {
	Parameter      string
	ParameterValue string
}

func (e *InvalidParameterValueError) Error() string {
	return fmt.Sprintf(`invalid value "%s" for parameter "%s"`, e.ParameterValue, e.Parameter) + "\r\n" + HowToUseTxt
}

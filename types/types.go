package types

import "fmt"

type CommandType string

const (
	CommandView CommandType = "view"
	CommandHelp CommandType = "help"
	CommandNone CommandType = "none"
)

func StringToCommandType(s string) (CommandType, error) {
	switch s {
	case string(CommandView):
		return CommandView, nil
	case string(CommandHelp):
		return CommandHelp, nil
	default:
		return CommandNone, fmt.Errorf("unknown command type: %s", s)
	}
}

package types

import "fmt"

type CommandType string

const (
	// public commands
	CommandHelp   CommandType = "help"
	CommandStart  CommandType = "start"  // send request to admin to confirm the user
	CommandDelete CommandType = "delete" // will delete the user from DB

	// private commands for confirmed users
	CommandView     CommandType = "view" // view project
	CommandRegister CommandType = "register"
	CommandNone     CommandType = "none"

	// admin commands
	CommandConfirm      CommandType = "confirm"
	CommandRequestsList CommandType = "requests_list" // list request for register
	CommandUsersList    CommandType = "users_list"
	CommandRemove       CommandType = "remove"
)

func StringToCommandType(s string) (CommandType, error) {
	switch s {
	case string(CommandView):
		return CommandView, nil
	case string(CommandHelp):
		return CommandHelp, nil
	case string(CommandStart):
		return CommandStart, nil
	case string(CommandDelete):
		return CommandDelete, nil
	case string(CommandRegister):
		return CommandRegister, nil
	case string(CommandConfirm):
		return CommandConfirm, nil
	case string(CommandRequestsList):
		return CommandRequestsList, nil
	case string(CommandUsersList):
		return CommandUsersList, nil
	case string(CommandRemove):
		return CommandRemove, nil
	default:
		return CommandNone, fmt.Errorf("unknown command type: %s", s)
	}
}

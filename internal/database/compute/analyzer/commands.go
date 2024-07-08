package analyzer

const (
	commandUnknownID = iota
	commandSetID
	commandGetID
	commandDelID
)

const (
	commandUnknown = "UNKNOWN"
	commandSet     = "SET"
	commandGet     = "GET"
	commandDel     = "DEL"
)

var commandNamesToID = map[string]int{
	commandUnknown: commandUnknownID,
	commandSet:     commandSetID,
	commandGet:     commandGetID,
	commandDel:     commandDelID,
}

func getCommandIDByName(commandName string) int {
	return commandNamesToID[commandName]
}

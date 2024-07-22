package analyzer

const (
	CommandUnknownID = iota
	CommandSetID
	CommandGetID
	CommandDelID
)

const (
	CommandUnknown = "UNKNOWN"
	CommandSet     = "SET"
	CommandGet     = "GET"
	CommandDel     = "DEL"
)

var commandNamesToID = map[string]int{
	CommandUnknown: CommandUnknownID,
	CommandSet:     CommandSetID,
	CommandGet:     CommandGetID,
	CommandDel:     CommandDelID,
}

func getCommandIDByName(commandName string) int {
	return commandNamesToID[commandName]
}

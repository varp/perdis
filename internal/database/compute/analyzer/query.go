package analyzer

type Query struct {
	commandId int
	args      []string
}

func NewQuery(commandId int, args []string) *Query {
	return &Query{
		commandId: commandId,
		args:      args,
	}
}

func (q *Query) CommandId() int {
	return q.commandId
}

func (q *Query) Args() []string {
	return q.args
}

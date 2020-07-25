package listcommands

import (
	"github.com/adlio/trello"
)

type ListCommand interface {
	UpdateList(list *trello.List) error
}

func ProcessList(list *trello.List, command ListCommand, c chan error) {
	c <- command.UpdateList(list)
}

func AllCommands() []ListCommand {
	return []ListCommand{
		SortCommand{},
	}
}

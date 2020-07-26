package cardcommands

import (
	"github.com/adlio/trello"
)

type CardCommand interface {
	CheckCard(card *trello.Card) bool
	UpdateCard(card *trello.Card) error
}

func ProcessCard(card *trello.Card, command CardCommand, c chan error) {
	if !card.Closed && command.CheckCard(card) {
		c <- command.UpdateCard(card)
	} else {
		c <- nil
	}
}

func AllCommands(members []string, delays map[string]float64) []CardCommand {
	return []CardCommand{
		MissingDueCommand{Delays: delays},
		MissingMemberCommand{Members: members},
		OverdueCommand{Delays: delays},
		MoveDueCommand{Delays: delays},
	}
}

package cardcommands

import "github.com/adlio/trello"

type OverdueCommand struct {
	Delays map[string]float64
}

func (OverdueCommand) CheckCard(card *trello.Card) bool {
	return card.DueComplete
}

func (OverdueCommand) UpdateCard(card *trello.Card) error {
	return nil
}

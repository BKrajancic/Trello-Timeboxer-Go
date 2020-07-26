package cardcommands

import (
	"strings"
	"time"

	"github.com/adlio/trello"
)

type MissingDueCommand struct {
	Delays map[string]float64
}

func (MissingDueCommand) CheckCard(card *trello.Card) bool {
	return card.Due == nil
}

func (m MissingDueCommand) UpdateCard(card *trello.Card) error {
	for match, days := range m.Delays {
		extraHours := time.Hour * time.Duration((24 * days))

		if strings.Contains(strings.ToLower(card.Name), strings.ToLower(match)) {
			due := time.Now().Add(extraHours)
			return card.Update(trello.Arguments{"due": due.String()})
		}
	}

	for match, days := range m.Delays {
		extraHours := time.Hour * time.Duration((24 * days))

		if strings.Contains(strings.ToLower(card.List.Name), strings.ToLower(match)) {
			due := time.Now().Add(extraHours)
			return card.Update(trello.Arguments{"due": due.Format(time.RFC3339)})
		}
	}

	return nil // Throw an error! Why are there no matches?!
}

package listcommands

import (
	"sort"

	"github.com/adlio/trello"
)

type SortCommand struct {
}

func (m SortCommand) UpdateList(list *trello.List) error {
	cards, err := list.GetCards(trello.Defaults())
	if err != nil {
		return err
	}

	// Don't need to sort a list of just one card.
	// We can now assume at least two cards for the rest of this function.
	if len(cards) < 2 {
		return nil
	}

	filtered_cards := make([]*trello.Card, 0)
	sorted := true
	prev_due := cards[0].Due

	for _, card := range cards {
		if !card.Closed && card.Due != nil {
			filtered_cards = append(filtered_cards, card)
			if prev_due != nil && card.Due.Before(*prev_due) {
				sorted = false
			}
			prev_due = card.Due
		}
	}

	if sorted {
		return nil
	}

	sort.Slice(filtered_cards, func(i, j int) bool {
		return filtered_cards[i].Due.Before(*filtered_cards[j].Due)
	})

	for i, card := range filtered_cards {
		err = card.SetPos(float64(i + 1.0))
		if err != nil {
			return err
		}
	}

	return err
}

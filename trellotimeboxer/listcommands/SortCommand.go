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

	sort.Slice(cards, func(i, j int) bool {
		if cards[i].Due == nil && cards[j] != nil {
			return true
		} else if cards[j].Due == nil && cards[i].Due != nil {
			return false
		}
		return cards[i].Due.Before(*cards[j].Due)
	})

	for _, card := range cards {
		if !card.Closed {
			err = card.MoveToBottomOfList()
			if err != nil {
				return err
			}
		}
	}

	return err
}

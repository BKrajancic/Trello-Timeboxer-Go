package cardcommands

import (
	"sort"
	"strings"
	"time"

	"github.com/adlio/trello"
)

type MoveDueCommand struct {
	Delays map[string]float64
}

func (MoveDueCommand) CheckCard(card *trello.Card) bool {
	return true // May as well do work in UpdateCard anyways.
}

func (m MoveDueCommand) UpdateCard(card *trello.Card) error {
	// Doesn't work sort by smallest to largest days
	invertedDelays := make(map[time.Time]string)
	maxDuesSorted := make([]time.Time, len(m.Delays))

	i := 0

	for match, days := range m.Delays {
		maxDue := time.Now().Add(time.Hour * time.Duration((24.0 * days)))
		invertedDelays[maxDue] = match
		maxDuesSorted[i] = maxDue
		i++
	}

	sort.Slice(maxDuesSorted, func(i, j int) bool {
		return maxDuesSorted[i].Before(maxDuesSorted[j])
	})

	for _, maxDue := range maxDuesSorted {
		if card.Due != nil && card.Due.Before(maxDue) {
			board := card.Board
			lists, err := board.GetLists(trello.Defaults())

			if err != nil {
				return err
			}

			match := invertedDelays[maxDue]

			for _, list := range lists {
				if strings.Contains(strings.ToLower(list.Name), strings.ToLower(match)) {
					if list.ID == card.IDList {
						return nil
					}
					return card.MoveToList(list.ID, trello.Defaults())
				}
			}
		}
	}

	return nil
}

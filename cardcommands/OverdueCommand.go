package cardcommands

import (
	"sort"
	"strings"
	"time"

	"github.com/adlio/trello"
)

type OverdueCommand struct {
	Delays map[string]float64
}

func (OverdueCommand) CheckCard(card *trello.Card) bool {
	return card.Due != nil && time.Now().After(*card.Due)
}

func (m OverdueCommand) UpdateCard(card *trello.Card) error {
	// Doesn't work sort by smallest to largest days
	invertedDelays := make(map[time.Time]string)
	maxDuesSorted := make([]time.Time, len(m.Delays))

	i := 0

	for match, days := range m.Delays {
		maxDue := time.Now().Add(-time.Hour * time.Duration((24.0 * 2 * days)))
		invertedDelays[maxDue] = match
		maxDuesSorted[i] = maxDue
		i++
	}

	sort.Slice(maxDuesSorted, func(i, j int) bool {
		return maxDuesSorted[i].After(maxDuesSorted[j])
	})

	for _, maxDue := range maxDuesSorted {
		// Filter only greater
		if card.Due.After(maxDue) {
			board := card.Board
			lists, err := board.GetLists(trello.Defaults())

			if err != nil {
				return err
			}

			match := invertedDelays[maxDue]

			for _, list := range lists {
				if strings.Contains(strings.ToLower(list.Name), strings.ToLower(match)) {
					if list.ID == card.IDList {
						if card.Due.After(maxDue) {
							return nil
						}
					} else {
						err := card.Update(trello.Arguments{"due": ""})
						if err != nil {
							return err
						}
						return card.MoveToList(list.ID, trello.Defaults())
					}
				}
			}
		}
	}

	return nil
}

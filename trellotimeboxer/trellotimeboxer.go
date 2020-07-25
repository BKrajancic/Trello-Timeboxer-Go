package trellotimeboxer

import (
	"fmt"
	"trellotimeboxer/cardcommands"
	"trellotimeboxer/config"
	"trellotimeboxer/listcommands"

	"github.com/adlio/trello"
)

func main() {
	trellotimeboxer()
}

func trellotimeboxer() error {
	config, err := config.GetConfig()

	if err == nil {
		client := trello.NewClient(config.AppKey, config.Token)

		board, err := client.GetBoard(config.BoardID, trello.Defaults())
		if err == nil {
			c := make(chan error)
			defer close(c)

			go processLists(board, listcommands.AllCommands(), c)
			err = <-c
			if err != nil {
				return nil
			}

			go processCards(board, cardcommands.AllCommands(config.Members, config.Delays), c)
			err = <-c
			if err != nil {
				return nil
			}
		}
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func processCards(board *trello.Board, commands []cardcommands.CardCommand, c chan error) {
	cards, err := board.GetCards(trello.Defaults())

	if err != nil {
		handleError(err)
	}

	numResults := len(commands) * len(cards)

	c2 := make(chan error, numResults)
	defer close(c2)

	for _, command := range commands {
		for _, card := range cards {
			card.Board = board
			go cardcommands.ProcessCard(card, command, c2)
		}
	}

	var out error = nil
	for i := 0; out == nil && i < numResults; i++ {
		out = <-c2
	}
	c <- out
}

func processLists(board *trello.Board, commands []listcommands.ListCommand, c chan error) {
	lists, err := board.GetLists(trello.Defaults())

	if err != nil {
		handleError(err)
	}

	numResults := len(commands) * len(lists)

	c2 := make(chan error, numResults)
	defer close(c2)

	for _, command := range commands {
		for _, list := range lists {
			go listcommands.ProcessList(list, command, c)
		}
	}

	var out error = nil
	for i := 0; out == nil && i < numResults; i++ {
		out = <-c2
	}
	c <- out
}

func handleError(err error) {
	fmt.Println(err.Error())
}

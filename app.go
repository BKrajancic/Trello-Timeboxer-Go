package main

import (
	"fmt"

	"github.com/BKrajancic/trellotimeboxergo/cardcommands"
	"github.com/BKrajancic/trellotimeboxergo/config"
	"github.com/BKrajancic/trellotimeboxergo/listcommands"

	"github.com/adlio/trello"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func run() error {
	config, err := config.GetConfig()

	if err != nil {
		return err
	}

	client := trello.NewClient(config.AppKey, config.Token)
	board, err := client.GetBoard(config.BoardID, trello.Defaults())
	if err != nil {
		return err
	}

	c := make(chan error)
	defer close(c)

	go processCards(board, cardcommands.AllCommands(config.Members, config.Delays), c)
	if err = <-c; err != nil {
		return err
	}

	go processLists(board, listcommands.AllCommands(), c)

	if err = <-c; err != nil {
		return err
	}

	return err
}

func processCards(board *trello.Board, commands []cardcommands.CardCommand, c chan error) {
	allCards := []*trello.Card{}

	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		fmt.Println(err.Error())
		c <- err
	}

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err == nil {
			for _, card := range cards {
				card.Board = board
				card.List = list
				allCards = append(allCards, card)
			}
		}
	}

	numResults := len(allCards) * len(commands)
	c2 := make(chan error, numResults)
	defer close(c2)
	for _, card := range allCards {
		for _, command := range commands {
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
		c <- err
		return
	}

	numResults := len(commands) * len(lists)

	c2 := make(chan error, numResults)
	defer close(c2)

	for _, command := range commands {
		for _, list := range lists {
			go listcommands.ProcessList(list, command, c2)
		}
	}

	var out error = nil
	for i := 0; out == nil && i < numResults; i++ {
		out = <-c2
	}
	c <- out
}

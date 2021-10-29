package main

import (
	"github.com/aneshas/holdem/deck"
	"github.com/aneshas/holdem/game"
)

var g *game.Game

type GameState struct {
	Players   []*game.Player
	Flop      game.Flop
	Turn      *deck.Card
	River     *deck.Card
	Hands     game.Hands
	Blinds    []game.PlayerID
	IsStarted bool
}

func mapGameState(g *game.Game) GameState {
	return GameState{
		Players:   g.Players(),
		Flop:      g.Flop(),
		Turn:      g.Turn(),
		River:     g.River(),
		Hands:     g.Hands(),
		IsStarted: g.Started,
		Blinds:    g.Blinds(),
	}
}

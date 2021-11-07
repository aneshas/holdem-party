package main

import "github.com/aneshas/holdem/game"

var store GameStore = GameStore{
	mem: make(map[game.ID]*game.Game),
}

type GameStore struct {
	mem map[game.ID]*game.Game
}

func (s *GameStore) Find(id game.ID) *game.Game {
	return s.mem[id]
}

func (s *GameStore) Save(game *game.Game) {
	s.mem[game.ID] = game
}

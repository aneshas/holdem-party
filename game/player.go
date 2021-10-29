package game

import "github.com/google/uuid"

type PlayerID string

func NewPlayer() *Player {
	return &Player{
		ID: PlayerID(uuid.NewString()),
	}
}

type Player struct {
	ID PlayerID
}

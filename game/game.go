package game

import (
	"errors"

	"github.com/aneshas/holdem/deck"
	"github.com/google/uuid"
)

var (
	ErrNoMinimumNumberOfPlayers      = errors.New("at least two players needed to start a game")
	ErrMaximumNumberOfPlayersReached = errors.New("maximum of nine players supported")
	ErrGameNotStarted                = errors.New("game should be started for this action")
	ErrGameAlreadyStarted            = errors.New("a game is already in progress")
)

func New() *Game {
	return &Game{
		ID:    ID(uuid.NewString()),
		hands: make(Hands),
	}
}

type ID string

type Game struct {
	ID          ID
	Started     bool
	players     []*Player
	hands       Hands
	deck        *deck.Deck
	flop        Flop
	turn        *deck.Card
	river       *deck.Card
	blinds      []PlayerID
	playersLeft bool
}

type Hand []deck.Card

type Hands map[PlayerID]Hand

type Flop []deck.Card

func (g *Game) Start() error {
	if g.Started && (g.flop == nil || g.turn == nil || g.river == nil) {
		return ErrGameAlreadyStarted
	}

	if len(g.players) < 2 {
		return ErrNoMinimumNumberOfPlayers
	}

	g.deck = deck.New()
	g.setBlinds()
	g.dealHands()
	g.Started = true
	g.flop = nil
	g.turn = nil
	g.river = nil
	g.playersLeft = false

	return nil
}

func (g *Game) setBlinds() {
	if g.playersLeft {
		g.blinds = nil
	}

	if g.blinds == nil || len(g.blinds) == 0 {
		g.blinds = []PlayerID{g.players[0].ID, g.players[1].ID}

		return
	}

	small, big := g.blinds[0], g.blinds[1]

	g.setBlind(0, small)
	g.setBlind(1, big)
}

func (g *Game) setBlind(i int, b PlayerID) {
	found := false

outer:
	for {
		for _, p := range g.players {
			if found {
				g.blinds[i] = p.ID
				break outer
			}

			if p.ID == b {
				found = true
				continue
			}
		}
	}
}

func (g *Game) dealHands() {
	for _, p := range g.players {
		g.hands[p.ID] = []deck.Card{}
	}

	for i := 0; i < 2; i++ {
		for _, p := range g.players {
			card, _ := g.deck.DrawTop()
			g.hands[p.ID] = append(g.hands[p.ID], card)
		}
	}
}

func (g *Game) Join(p *Player) (int, error) {
	if len(g.players) == 9 {
		return 0, ErrMaximumNumberOfPlayersReached
	}

	g.players = append(g.players, p)

	return len(g.players), nil
}

func (g *Game) Hands() Hands {
	if g.Started {
		return nil
	}

	return g.hands
}

func (g *Game) DealNext() error {
	if !g.Started {
		return ErrGameNotStarted
	}

	switch {
	case len(g.flop) != 3:
		g.dealFlop()
	case g.turn == nil:
		g.dealTurn()
	case g.river == nil:
		g.dealRiver()
	default:
		g.Started = false
	}

	return nil
}

func (g *Game) dealFlop() {
	g.deck.DrawTop()

	f1, _ := g.deck.DrawTop()
	f2, _ := g.deck.DrawTop()
	f3, _ := g.deck.DrawTop()

	g.flop = []deck.Card{
		f1, f2, f3,
	}
}

func (g *Game) dealTurn() {
	g.deck.DrawTop()

	turn, _ := g.deck.DrawTop()

	g.turn = &turn
}

func (g *Game) dealRiver() {
	g.deck.DrawTop()

	river, _ := g.deck.DrawTop()

	g.river = &river
}

func (g *Game) Flop() Flop { return g.flop }

func (g *Game) Turn() *deck.Card { return g.turn }

func (g *Game) River() *deck.Card { return g.river }

func (g *Game) Fold(id PlayerID) {
	delete(g.hands, id)

	if g.Started && len(g.hands) == 0 {
		g.DealNext()
		g.DealNext()
		g.DealNext()
		g.Start()
	}
}

func (g *Game) Players() []*Player { return g.players }

func (g *Game) PlayerHand(id PlayerID) Hand {
	if h, ok := g.hands[id]; ok {
		return h
	}

	return nil
}

func (g *Game) PlayerSeatNumber(id PlayerID) int {
	for i, p := range g.players {
		if p.ID == id {
			return i + 1
		}
	}

	return 0
}

func (g *Game) Blinds() []PlayerID { return g.blinds }

func (g *Game) Leave(id PlayerID) {
	g.Fold(id)

	var newPlayers []*Player

	for _, p := range g.players {
		if id == p.ID {
			continue
		}

		newPlayers = append(newPlayers, p)
	}

	g.players = newPlayers
	g.playersLeft = true
}

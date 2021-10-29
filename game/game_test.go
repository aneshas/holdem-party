package game_test

import (
	"testing"

	"github.com/aneshas/holdem/game"
)

func TestCantStartAGameWithNoPlayers(t *testing.T) {
	g := game.New()

	err := g.Start()
	if err != game.ErrNoMinimumNumberOfPlayers {
		t.Fatal("should not be able to start a game with no players")
	}
}

func TestCantStartAGameWithOnePlayer(t *testing.T) {
	g := game.New()

	g.Join(game.NewPlayer())

	err := g.Start()
	if err != game.ErrNoMinimumNumberOfPlayers {
		t.Fatal("should not be able to start a game with no players")
	}
}

func TestOnlyNinePlayersCanJoin(t *testing.T) {
	g := game.New()

	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())

	_, err := g.Join(game.NewPlayer())
	if err != game.ErrMaximumNumberOfPlayersReached {
		t.Fatal("should not be able to join more than nine players")
	}
}

func TestGameCanBeStartedWithMinimumNumberOfPlayers(t *testing.T) {
	g := game.New()

	g.Join(game.NewPlayer())
	g.Join(game.NewPlayer())

	err := g.Start()
	if err != nil {
		t.Fatal("should be able to start the game")
	}
}

func TestOnlyOneGameCanBeInProgress(t *testing.T) {
	g := newGameWithTwoPlayers()

	g.Start()
	err := g.Start()

	if err != game.ErrGameAlreadyStarted {
		t.Fatal("should not be possible to start a game while a game is in progress")
	}
}

func TestHandsNotAvailableWhileGameInProgress(t *testing.T) {
	g := game.New()

	p1 := game.NewPlayer()
	p2 := game.NewPlayer()

	g.Join(p1)
	g.Join(p2)

	g.Start()

	if g.Hands() != nil {
		t.Fatal("hands should not be available while game is in progress")
	}
}

func TestHandsShouldBeAvailableAfterGameIsFinished(t *testing.T) {
	g := game.New()

	p1 := game.NewPlayer()
	p2 := game.NewPlayer()

	g.Join(p1)
	g.Join(p2)

	g.Start()

	g.DealNext()
	g.DealNext()
	g.DealNext()
	g.DealNext()

	if len(g.Hands()[p1.ID]) != 2 {
		t.Fatal("player one should have been dealt two cards")
	}

	if len(g.Hands()[p2.ID]) != 2 {
		t.Fatal("player two should have been dealt two cards")
	}
}

func TestShouldOnlyGetHandsForNonFoldedPlayers(t *testing.T) {
	g := game.New()

	p1 := game.NewPlayer()
	p2 := game.NewPlayer()

	g.Join(p1)
	g.Join(p2)

	g.Start()

	g.DealNext()
	g.DealNext()
	g.Fold(p1.ID)
	g.DealNext()
	g.DealNext()

	if _, ok := g.Hands()[p1.ID]; ok || g.Hands() == nil {
		t.Fatal("folded hands should not be shown")
	}
}

func TestFlopShouldBeDealt(t *testing.T) {
	g := newGameWithTwoPlayers()

	g.Start()
	g.DealNext()

	if len(g.Flop()) != 3 {
		t.Fatal("three flop cards should have been dealt")
	}

	if g.Turn() != nil {
		t.Fatal("turn should have not been dealt at this time")
	}

	if g.River() != nil {
		t.Fatal("river should have not been dealt at this time")
	}
}

func TestCantDealFlopIfGameNotStarted(t *testing.T) {
	g := newGameWithTwoPlayers()

	err := g.DealNext()

	if err != game.ErrGameNotStarted {
		t.Fatal("dealing flop should not be possible if game is not started")
	}
}

func TestTurnShouldBeDealt(t *testing.T) {
	g := newGameWithTwoPlayers()

	g.Start()
	g.DealNext()
	g.DealNext()

	if g.Turn() == nil {
		t.Fatal("turn should have been dealt")
	}

	if g.River() != nil {
		t.Fatal("river should have not been dealt at this time")
	}
}

func TestRiverShouldBeDealt(t *testing.T) {
	g := newGameWithTwoPlayers()

	g.Start()
	g.DealNext()
	g.DealNext()
	g.DealNext()

	if g.River() == nil {
		t.Fatal("river should have been dealt")
	}
}

func TestGameCanBeReStartedAfterRiver(t *testing.T) {
	g := newGameWithTwoPlayers()

	g.Start()
	g.DealNext()
	g.DealNext()
	g.DealNext()

	err := g.Start()
	if err != nil {
		t.Fatal("should be able to restart the game after river")
	}

	if g.Flop() != nil || g.Turn() != nil || g.River() != nil {
		t.Fatal("flop, turn and river should have been reset")
	}
}

func TestBlindsRotate(t *testing.T) {
	g := game.New()

	p1 := game.NewPlayer()
	p2 := game.NewPlayer()

	g.Join(p1)
	g.Join(p2)

	g.Start()

	blinds := g.Blinds()

	if blinds[0] != p1.ID {
		t.Fatal("player one should have been the small blind")
	}

	if blinds[1] != p2.ID {
		t.Fatal("player two should have been the big blind")
	}

	g.DealNext()
	g.DealNext()
	g.DealNext()

	g.Start()

	if blinds[0] != p2.ID {
		t.Fatal("player two should have been the small blind")
	}

	if blinds[1] != p1.ID {
		t.Fatal("player one should have been the big blind")
	}
}

func newGameWithTwoPlayers() *game.Game {
	g := game.New()

	p1 := game.NewPlayer()
	p2 := game.NewPlayer()

	g.Join(p1)
	g.Join(p2)

	return g
}

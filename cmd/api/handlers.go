package main

import (
	"net/http"

	"github.com/aneshas/holdem/deck"
	"github.com/aneshas/holdem/game"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string
}

func handleGame(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	state := mapGameState(g)

	c.JSON(http.StatusOK, state)
}

type JoinedPlayer struct {
	PlayerNumber int
	PlayerID     string
}

func handleJoin(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	p := game.NewPlayer()

	num, err := g.Join(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, JoinedPlayer{
		PlayerNumber: num,
		PlayerID:     string(p.ID),
	})
}

func handleStart(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	err := g.Start()
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapGameState(g))
}

func handleProceed(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	err := g.DealNext()
	if err != nil {
		c.JSON(http.StatusBadRequest, Error{
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, mapGameState(g))
}

type PlayerSession struct {
	Hand       []deck.Card
	SeatNumber int
}

func handlePlayerSession(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	id := game.PlayerID(c.Param("pid"))

	c.JSON(http.StatusOK, PlayerSession{
		Hand:       g.PlayerHand(id),
		SeatNumber: g.PlayerSeatNumber(id),
	})
}

func handlePlayerFold(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	g.Fold(game.PlayerID(c.Param("pid")))

	c.Status(http.StatusOK)
}

func handleNewGame(c *gin.Context) {
	g := game.New()

	store.Save(g)

	c.JSON(http.StatusOK, mapGameState(g))
}

func handlePlayerLeave(c *gin.Context) {
	g := store.Find(game.ID(c.Param("id")))

	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	g.Leave(game.PlayerID(c.Param("pid")))

	c.Status(http.StatusOK)
}

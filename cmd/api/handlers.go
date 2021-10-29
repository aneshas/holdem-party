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
	createGame()

	state := mapGameState(g)

	c.JSON(http.StatusOK, state)
}

type JoinedPlayer struct {
	PlayerNumber int
	PlayerID     string
}

func handleJoin(c *gin.Context) {
	createGame()

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
	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

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
	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

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
	Hand []deck.Card
}

func handlePlayerSession(c *gin.Context) {
	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	hand := g.PlayerHand(game.PlayerID(c.Param("id")))

	c.JSON(http.StatusOK, PlayerSession{
		Hand: hand,
	})
}

func handlePlayerFold(c *gin.Context) {
	if g == nil {
		c.Status(http.StatusBadRequest)
		return
	}

	g.Fold(game.PlayerID(c.Param("id")))

	c.Status(http.StatusOK)
}

func handleNewGame(c *gin.Context) {
	g = game.New()
	c.Status(http.StatusOK)
}

func createGame() {
	if g == nil {
		g = game.New()
	}
}

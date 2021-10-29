package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.GET("/game", handleGame)
	r.GET("/game/new", handleNewGame)
	r.GET("/game/join", handleJoin)
	r.GET("/game/start", handleStart)
	r.GET("/game/proceed", handleProceed)
	r.GET("/game/player/:id", handlePlayerSession)
	r.GET("/game/player/:id/fold", handlePlayerFold)
	r.GET("/game/player/:id/leave", handlePlayerLeave)

	log.Fatal(r.Run(fmt.Sprintf("0.0.0.0:%d", 8080)))
}

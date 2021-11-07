package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.Static("/holdem", "./build")
	r.Static("/static", "./build/static")
	r.Static("/assets", "./build/assets")
	r.Static("/favicon.ico", "./build/favicon.ico")
	r.GET("/game/:id", handleGame)
	r.GET("/game/new", handleNewGame)
	r.GET("/game/:id/join", handleJoin)
	r.GET("/game/:id/start", handleStart)
	r.GET("/game/:id/proceed", handleProceed)
	r.GET("/game/:id/player/:pid", handlePlayerSession)
	r.GET("/game/:id/player/:pid/fold", handlePlayerFold)
	r.GET("/game/:id/player/:pid/leave", handlePlayerLeave)

	log.Fatal(r.Run(fmt.Sprintf("0.0.0.0:%d", 8080)))
}

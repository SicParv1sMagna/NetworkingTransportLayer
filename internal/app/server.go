package app

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *Application) StartServer() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/http")
	{
		api.POST("/send", a.handler.SendMessage)
		api.GET("/recieve")
		api.POST("/transfer")
	}

	err := router.Run()
	if err != nil {
		log.Println("Error with running server")
		return
	}
}

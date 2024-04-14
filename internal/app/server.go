package app

import (
	"log"

	"github.com/SicParv1sMagna/NetworkingTransportLayer/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Application) StartServer() {
	router := gin.Default()

	docs.SwaggerInfo.Title = "Транспортный уровень"
	docs.SwaggerInfo.Description = ""
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := router.Group("/http")
	{
		api.POST("/send", a.handler.SendMessage)
		api.POST("/transfer", a.handler.TransferSegments)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run()
	if err != nil {
		log.Println("Error with running server")
		return
	}
}

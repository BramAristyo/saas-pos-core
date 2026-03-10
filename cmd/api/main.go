package main

import (
	"log"

	handler "github.com/BramAristyo/go-pos-mawish/internal/handlers"
	"github.com/BramAristyo/go-pos-mawish/internal/infra/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/internal/repositories"
	"github.com/BramAristyo/go-pos-mawish/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitDb()
	defer database.CloseDb()

	if err != nil {
		log.Fatal(err.Error())
	}

	db := database.GetDb()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/api")
	{
		v1.GET("/", userHandler.GetAll)
	}

	router.Run(":9000")
}

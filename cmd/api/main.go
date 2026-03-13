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

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	{
		users.GET("/", userHandler.GetAll)
		users.GET("/:id", userHandler.FindById)
		users.POST("/", userHandler.Store)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	categories := v1.Group("/categories")
	{
		categories.GET("/", categoryHandler.Paginate)
		categories.GET("/:id", categoryHandler.FindById)
		categories.POST("/", categoryHandler.Store)
		categories.PUT("/:id", categoryHandler.Update)
	}

	router.Run(":9000")
}

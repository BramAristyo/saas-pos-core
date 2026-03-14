package main

import (
	"log"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/internal/config"
	"github.com/BramAristyo/go-pos-mawish/internal/handler"
	"github.com/BramAristyo/go-pos-mawish/internal/infra/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/internal/middleware"
	"github.com/BramAristyo/go-pos-mawish/internal/repository"
	"github.com/BramAristyo/go-pos-mawish/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	err := database.InitDb(cfg)
	defer database.CloseDb()

	if err != nil {
		log.Fatal(err.Error())
	}

	db := database.GetDb()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router := gin.Default()
	s := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}
	router.Use(middleware.ErrorHandler())

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

	s.ListenAndServe()
}

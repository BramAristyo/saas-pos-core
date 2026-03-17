package main

import (
	"log"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/internal/config"
	"github.com/BramAristyo/go-pos-mawish/internal/dependecy"
	"github.com/BramAristyo/go-pos-mawish/internal/infra/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/internal/middleware"
	"github.com/BramAristyo/go-pos-mawish/internal/router"
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

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	handlers := dependecy.Bootstrap(db, cfg)
	router.RegisterRoutes(r, handlers, cfg)

	s := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	s.ListenAndServe()
}

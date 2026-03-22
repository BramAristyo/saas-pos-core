package main

import (
	"log"
	"net/http"

	"github.com/BramAristyo/go-pos-mawish/internal/api/middleware"
	"github.com/BramAristyo/go-pos-mawish/internal/api/router"
	"github.com/BramAristyo/go-pos-mawish/internal/dependency"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
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

	handlers := dependency.Bootstrap(db, cfg)
	router.RegisterRoutes(r, handlers, cfg)

	s := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Http server failed: %v", err)
	}
}

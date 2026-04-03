package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BramAristyo/go-pos-mawish/internal/api/middleware"
	"github.com/BramAristyo/go-pos-mawish/internal/api/router"
	"github.com/BramAristyo/go-pos-mawish/internal/dependency"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
	"github.com/BramAristyo/go-pos-mawish/pkg/logger"
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: Implement Dashboard API One API for All use Goroutines! (Business)
	cfg := config.GetConfig()
	zapLogger := logger.NewZapLogger(cfg)

	err := database.InitDb(cfg, zapLogger)
	defer database.CloseDb()

	if err != nil {
		zapLogger.Fatal(err.Error())
	}

	db := database.GetDb()

	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RateLimiter(zapLogger))

	r.Use(middleware.ErrorHandler(zapLogger))

	r.Use(ginzap.Ginzap(zapLogger.GetLogger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zapLogger.GetLogger(), true))

	handlers := dependency.Bootstrap(db, cfg)
	router.RegisterRoutes(r, handlers, cfg)

	s := &http.Server{
		Addr:         cfg.Server.Addr,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zapLogger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	zapLogger.Info("Server existing")
}

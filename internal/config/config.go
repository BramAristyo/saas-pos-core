package config

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Addr           string
	Handler        *gin.Engine
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DbName          string
	SSLMode         string
	MaxIdleCons     int
	MaxOpenCons     int
	ConnMaxLifetime time.Duration
}

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

func GetConfig() *Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &Config{
		Server: ServerConfig{
			Addr:           viper.GetString("SERVER_ADDR"),
			ReadTimeout:    viper.GetDuration("SERVER_READ_TIMEOUT"),
			WriteTimeout:   viper.GetDuration("SERVER_WRITE_TIMEOUT"),
			MaxHeaderBytes: viper.GetInt("SERVER_MAX_HEADER_BYTES"),
		},
		Postgres: PostgresConfig{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetString("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			DbName:          viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("SSL_MODE"),
			MaxIdleCons:     viper.GetInt("DB_MAX_IDLE_CONNS"),
			MaxOpenCons:     viper.GetInt("DB_MAX_OPEN_CONNS"),
			ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		},
	}
}

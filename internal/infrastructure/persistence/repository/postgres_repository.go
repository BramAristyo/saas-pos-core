package repository

import (
	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/persistence/database"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	DB       *gorm.DB
	Preloads database.PreloadEntity
}

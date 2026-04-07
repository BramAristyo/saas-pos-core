package repository

import (
	"github.com/BramAristyo/saas-pos-core/server/internal/infrastructure/persistence/database"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	DB       *gorm.DB
	Preloads database.PreloadEntity
}

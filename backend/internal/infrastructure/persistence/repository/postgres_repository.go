package repository

import (
	"github.com/BramAristyo/saas-pos-core/backend/internal/infrastructure/persistence/database"
	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	DB       *gorm.DB
	Preloads database.PreloadEntity
}

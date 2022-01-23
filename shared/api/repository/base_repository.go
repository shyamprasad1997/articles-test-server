package repository

import (
	"articles-test-server/infrastructure"
)

// BaseRepository struct.
type BaseRepository struct {
	Logger infrastructure.LoggerInterface
}

// NewBaseRepository returns NewBaseRepository instance.
func NewBaseRepository(logger infrastructure.LoggerInterface) *BaseRepository {
	return &BaseRepository{Logger: logger}
}

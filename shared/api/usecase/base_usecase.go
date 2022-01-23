package usecase

import (
	"articles-test-server/infrastructure"
)

// BaseRepository struct.
type BaseUsecase struct {
	Logger infrastructure.LoggerInterface
}

// NewBaseRepository returns NewBaseRepository instance.
func NewBaseUsecase(logger infrastructure.LoggerInterface) *BaseUsecase {
	return &BaseUsecase{Logger: logger}
}

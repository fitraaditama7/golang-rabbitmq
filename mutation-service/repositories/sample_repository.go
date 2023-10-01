package repository

import (
	"context"
	"mutation-service/entities"
)

//go:generate mockgen -destination=../mocks/repository/sample_repository_mock.go -package=mock_repository . SampleRepository
type SampleRepository interface {
	FindAll(ctx context.Context) ([]*entities.Sample, error)
	FindByID(ctx context.Context, id int64) (*entities.Sample, error)
	Create(ctx context.Context, sample *entities.Sample) error
	Update(ctx context.Context, sample *entities.Sample) error
	Delete(ctx context.Context, id string) (int64, error)
}

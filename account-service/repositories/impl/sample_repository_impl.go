package repository_impl

import (
	"account-service/entities"
	repository "account-service/repositories"
	"context"

	"gorm.io/gorm"
)

type SampleRepository struct {
	DB *gorm.DB
}

func NewSampleRepository(DB *gorm.DB) repository.SampleRepository {
	return &SampleRepository{DB: DB}
}

func (s SampleRepository) FindAll(ctx context.Context) ([]*entities.Sample, error) {
	var samples []*entities.Sample

	result := s.DB.WithContext(ctx).Find(&samples)
	if result.Error != nil {
		return nil, result.Error
	}

	return samples, nil
}

func (s SampleRepository) FindByID(ctx context.Context, id int64) (*entities.Sample, error) {
	var sample entities.Sample
	result := s.DB.WithContext(ctx).Where("id=?", id).First(&sample)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &sample, nil
}

func (s SampleRepository) Create(ctx context.Context, sample *entities.Sample) error {
	return s.DB.WithContext(ctx).Create(sample).Error
}

func (s SampleRepository) Update(ctx context.Context, sample *entities.Sample) error {
	return s.DB.WithContext(ctx).Save(&sample).Error
}

func (s SampleRepository) Delete(ctx context.Context, id string) (int64, error) {
	result := s.DB.WithContext(ctx).Delete(entities.Sample{}, "id = ?", id)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

package mapper

import (
	"account-service/entities"
	"account-service/models"
)

func SampleEntityToModel(sample *entities.Sample) models.SampleResponse {
	return models.SampleResponse{
		ID:        sample.ID,
		Name:      sample.Name,
		CreatedAt: sample.CreatedAt,
		UpdatedAt: sample.UpdatedAt,
	}
}

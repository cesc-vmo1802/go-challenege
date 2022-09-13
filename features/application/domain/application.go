package domain

import (
	"go-challenege/features/application/dto"
	"go-challenege/pkg/database"
)

const (
	Entity = "application"
)

type Application struct {
	database.MgoModel `bson:",inline"`
	Name              string  `bson:"name"`
	Description       string  `bson:"description"`
	Enabled           *bool   `bson:"enabled"`
	Type              *string `bson:"type"`
}

func FromDTO(input *dto.CreateApplicationRequest) Application {
	return Application{
		Name:        input.Name,
		Description: input.Description,
		Type:        &input.Type,
		Enabled:     &input.Enabled,
	}
}

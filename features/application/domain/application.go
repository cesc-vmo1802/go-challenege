package domain

import (
	"go-challenege/features/application/dto"
	"go-challenege/pkg/database"
)

const (
	Entity = "application"
)

type Application struct {
	database.MgoModel
	Name        string
	Description string
	Enabled     *bool
	Type        *string
}

func FromDTO(input *dto.CreateApplicationRequest) Application {
	return Application{
		Name:        input.Name,
		Description: input.Description,
	}
}

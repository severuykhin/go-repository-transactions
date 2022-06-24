package repository

import "main/internal/entities"

type AttributeRepository interface {
	Create(productId string, name string, value string) (*entities.Attribute, error)
}

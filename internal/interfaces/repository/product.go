package repository

import (
	"main/internal/entities"
)

type ProductRepository interface {
	Create(name string) (*entities.Product, error)
}

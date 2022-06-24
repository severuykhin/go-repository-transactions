package common

import "main/internal/interfaces/repository"

type RepoRegistry interface {
	GetProductRepo() repository.ProductRepository
	GetAttributeRepo() repository.AttributeRepository
}

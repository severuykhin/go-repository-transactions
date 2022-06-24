package registry

import (
	"main/internal/interfaces/common"
	"main/internal/interfaces/repository"
	"main/internal/repository/attribute"
	"main/internal/repository/product"
)

type pgRepoRegistry struct {
	client common.DBClient
}

func NewPgRepoRegistry(client common.DBClient) *pgRepoRegistry {
	return &pgRepoRegistry{
		client: client,
	}
}

func (r pgRepoRegistry) GetProductRepo() repository.ProductRepository {
	return product.NewProductPgRepository(r.client)
}

func (r pgRepoRegistry) GetAttributeRepo() repository.AttributeRepository {
	return attribute.NewAttributePgRepository(r.client)
}

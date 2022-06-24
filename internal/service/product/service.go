package product

import (
	"main/internal/entities"
	"main/internal/interfaces/common"
)

type productService struct {
	dbService common.DbService
}

func NewProductService(dbService common.DbService) *productService {
	return &productService{
		dbService: dbService,
	}
}

func (s productService) CreateProduct(name string, attributes map[string]string) (*entities.Product, error) {

	out, err := s.dbService.DoInTransaction(func(registry common.RepoRegistry) (interface{}, error) {

		productRepo := registry.GetProductRepo()
		attributesRepo := registry.GetAttributeRepo()

		product, err := productRepo.Create(name)

		if err != nil {
			return nil, err
		}

		for name, value := range attributes {
			attribute, err := attributesRepo.Create(product.Id, name, value)

			if err != nil {
				return nil, err
			}

			product.Attributes = append(product.Attributes, *attribute)
		}

		return product, nil
	})

	if err != nil {
		return nil, err
	}

	res := out.(*entities.Product)

	return res, nil
}

func (s productService) GetProduct(id string) (*entities.Product, error) {
	out, err := s.dbService.Do(func(registry common.RepoRegistry) (interface{}, error) {
		productRepo := registry.GetProductRepo()
		// attributesRepo := registry.GetAttributeRepo()

		product, err := productRepo.GetOne(id)
		if err != nil {
			return nil, err
		}

		return product, nil

	})

	if err != nil {
		return nil, err
	}

	res := out.(*entities.Product)
	return res, nil
}

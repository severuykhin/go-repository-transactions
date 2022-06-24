package product

import (
	"main/internal/entities"
	"main/internal/interfaces/common"
)

type productService struct {
	trxService common.TransactionService
}

func NewProductService(trxService common.TransactionService) *productService {
	return &productService{
		trxService: trxService,
	}
}

func (s productService) CreateProduct(name string, attributes map[string]string) (*entities.Product, error) {

	out, err := s.trxService.DoInTransaction(func(registry common.RepoRegistry) (interface{}, error) {

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

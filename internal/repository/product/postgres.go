package product

import (
	"main/internal/entities"
	"main/internal/interfaces/common"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
)

type productPgRepo struct {
	db common.DBClient
}

func NewProductPgRepository(client common.DBClient) *productPgRepo {
	return &productPgRepo{
		db: client,
	}
}

func (r *productPgRepo) Create(name string) (*entities.Product, error) {
	product := entities.Product{
		Id:   uuid.New().String(),
		Name: name,
	}

	query := sq.
		Insert("product").
		Columns(
			"id",
			"name",
		).
		Values(
			product.Id,
			product.Name,
		)

	sql, values, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sql = r.db.Rebind(sql)

	_, err = r.db.Exec(sql, values...)

	if err != nil {
		return nil, err
	}

	return &product, nil

}

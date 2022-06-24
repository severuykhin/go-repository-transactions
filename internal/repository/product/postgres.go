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

func (r *productPgRepo) GetOne(id string) (*entities.Product, error) {
	product := entities.Product{}

	query := sq.
		Select("*").
		From("product").
		Where(sq.Eq{
			"id": id,
		})

	sql, values, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	sql = r.db.Rebind(sql)

	res := r.db.QueryRow(sql, values...)
	err = res.Scan(&product.Id, &product.Name)

	if err != nil {
		return &product, err
	}

	return &product, nil
}

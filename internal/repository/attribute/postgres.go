package attribute

import (
	"fmt"
	"main/internal/entities"
	"main/internal/interfaces/common"

	sq "github.com/Masterminds/squirrel"
)

type attributePgRepo struct {
	db common.DBClient
}

func NewAttributePgRepository(client common.DBClient) *attributePgRepo {
	return &attributePgRepo{
		db: client,
	}
}

func (r *attributePgRepo) Create(productId string, name string, value string) (*entities.Attribute, error) {
	attribute := entities.Attribute{
		ProductId: productId,
		Name:      name,
		Value:     value,
	}

	query := sq.
		Insert("product_attribute").
		Columns(
			"product_id",
			"name",
			"value",
		).
		Values(
			attribute.ProductId,
			attribute.Name,
			attribute.Value,
		)

	sql, values, err := query.ToSql()
	if err != nil {
		fmt.Println("test")
		return nil, err
	}

	sql = r.db.Rebind(sql)

	_, err = r.db.Exec(sql, values...)

	if err != nil {
		return nil, err
	}

	return &attribute, nil

}

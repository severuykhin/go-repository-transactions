package transaction

import (
	"context"
	"database/sql"
	"main/internal/db/postgres"
	"main/internal/interfaces/common"
	"main/internal/repository/registry"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type PostgresTrasactionService struct {
	connection *sqlx.DB
}

func NewPostgresTransactionService(dbConnection *sqlx.DB) *PostgresTrasactionService {
	return &PostgresTrasactionService{
		connection: dbConnection,
	}
}

func (s *PostgresTrasactionService) DoInTransaction(trFunc common.TransactionClosure) (interface{}, error) {
	tx, err := s.connection.BeginTx(context.TODO(), nil)

	if err != nil {
		return nil, err
	}

	repoRegistry := registry.NewPgRepoRegistry(postgres.NewPgDbClient(s.connection, tx))

	res, err := trFunc(repoRegistry)

	if err != nil {
		return nil, completeTransaction(tx, err)
	}

	err = completeTransaction(tx, nil)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func completeTransaction(tx *sql.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Wrap(err, rollbackErr.Error())
		}
		return err
	} else {
		if commitErr := tx.Commit(); commitErr != nil {
			return commitErr
		}
		return nil
	}
}

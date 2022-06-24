package common

type TransactionClosure func(registry RepoRegistry) (interface{}, error)

type TransactionService interface {
	DoInTransaction(trFunc TransactionClosure) (interface{}, error)
}

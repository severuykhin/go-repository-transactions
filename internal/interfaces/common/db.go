package common

type TransactionClosure func(registry RepoRegistry) (interface{}, error)
type Closure func(registry RepoRegistry) (interface{}, error)

type DbService interface {
	DoInTransaction(trFunc TransactionClosure) (interface{}, error)
	Do(f Closure) (interface{}, error)
}

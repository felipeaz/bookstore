package database

type GORMServiceInterface interface {
	FetchAll(domainObj interface{}) (interface{}, error)
	FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, error)
	FetchAllWithQueryAndPreload(domainObj interface{}, query, preload, join, group string) (interface{}, error)
	Fetch(domainObj interface{}, id string) (interface{}, error)
	FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, error)
	FetchAllWhere(domainObj interface{}, query string) (interface{}, error)
	Persist(domainObj interface{}) error
	Refresh(domainObj interface{}, id string) error
	Remove(domainObj interface{}, id string) error
	RemoveWhere(domainObj interface{}, query string) error
	GetErrorStatusCode(err error) int
}

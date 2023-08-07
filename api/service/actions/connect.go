package actions

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"library/api/model"
)

//go:generate gmg --dst ./mocks/{}_gomock.go
type Connector interface {
	GetConnection() (*sql.DB, error)
}

type StorageConnector struct {
	options model.DBOptions
}

func NewStorageConnector(options model.DBOptions) *StorageConnector {
	return &StorageConnector{options: options}
}
func (c *StorageConnector) GetConnection() (*sql.DB, error) {
	//TODO:add authorization and authentication
	return sql.Open("mysql", c.options.User+":"+c.options.Password+"@tcp("+c.options.Connect+")/"+c.options.DBName)
}

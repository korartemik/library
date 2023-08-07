package storage

import (
	"context"
	"database/sql"
	"library/api/model"
)

//go:generate gmg --dst ./mocks/{}_gomock.go
type LibraryStorage interface {
	GetByTitle(ctx context.Context, title string, db *sql.DB) ([]*model.Book, error)
	GetByAuthor(ctx context.Context, title string, db *sql.DB) ([]*model.Book, error)
}

package library

import (
	"context"
	"database/sql"
	"fmt"
	model2 "library/api/model"
	dto2 "library/api/storage/dto"
)

type LibraryStorage struct {
}

func NewLibraryStorage() *LibraryStorage {
	return &LibraryStorage{}
}

func (s LibraryStorage) GetByTitle(ctx context.Context, title string, db *sql.DB) ([]*model2.Book, error) {
	var dtoBooks []*dto2.Book
	//TODO: add pagination
	rows, err := db.Query("select id, author, title from library where title = ?", title)
	if err != nil {
		return nil, fmt.Errorf("problem with query", err)
	}
	defer rows.Close()

	for rows.Next() {
		book := &dto2.Book{}
		err := rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, fmt.Errorf("problem with scan", err)
		}
		dtoBooks = append(dtoBooks, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(dtoBooks) == 0 {
		return nil, model2.NotFound(fmt.Errorf("Not found book by Title = '%s'", title))
	}

	var books []*model2.Book
	for _, book := range dtoBooks {
		modelBook, err := dto2.ConvertBookToModel(book)
		if err != nil {
			return nil, fmt.Errorf("problem with convert", err)
		}
		books = append(books, modelBook)
	}

	return books, nil
}

func (s LibraryStorage) GetByAuthor(ctx context.Context, author string, db *sql.DB) ([]*model2.Book, error) {
	var dtoBooks []*dto2.Book
	//TODO: add pagination
	rows, err := db.Query("select id, author, title from library where author = ?", author)
	if err != nil {
		return nil, fmt.Errorf("problem with query", err)
	}
	defer rows.Close()

	for rows.Next() {
		book := &dto2.Book{}
		err := rows.Scan(&book.ID, &book.Author, &book.Title)
		if err != nil {
			return nil, fmt.Errorf("problem with scan", err)
		}
		dtoBooks = append(dtoBooks, book)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(dtoBooks) == 0 {
		return nil, model2.NotFound(fmt.Errorf("Not found book by Author = '%s'", author))
	}

	var books []*model2.Book
	for _, book := range dtoBooks {
		modelBook, err := dto2.ConvertBookToModel(book)
		if err != nil {
			return nil, fmt.Errorf("problem with convert", err)
		}
		books = append(books, modelBook)
	}

	return books, nil
}

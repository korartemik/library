package dto

import (
	"library/api/model"
)

func ConvertBookToModel(from *Book) (*model.Book, error) {
	return &model.Book{
		Author: from.Author,
		Title:  from.Title,
	}, nil
}

func ConvertBookFromModel(from *model.Book) (*Book, error) {
	return &Book{
		Author: from.Author,
		Title:  from.Title,
	}, nil
}

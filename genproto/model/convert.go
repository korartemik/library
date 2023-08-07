package model

import (
	"library/api/model"
)

func ConvertToAPI(from *model.Book) *Book {
	return &Book{
		Title:  from.Title,
		Author: from.Author,
	}
}

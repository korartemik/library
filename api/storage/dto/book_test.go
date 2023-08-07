package dto

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"test-library/api/model"
	"testing"
)

func TestConvertBook(t *testing.T) {
	book := &model.Book{
		Author: "some-author",
		Title:  "some-name",
	}
	dtoBook, err := ConvertBookFromModel(book)
	require.NoError(t, err)
	gotBook, err := ConvertBookToModel(dtoBook)
	require.NoError(t, err)
	assert.Equal(t, book, gotBook)
}

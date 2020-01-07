package openapitools

import (
	"github.com/csandiego/poc-openapitools/go/client/data"
	"github.com/csandiego/poc-openapitools/go/openapitools"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreate(t *testing.T) {
	api := &testBookApiService{createErr: nil}
	service := NewBookService(api)
	book := data.Book{Title: "abc", Author: "xyz"}
	err := service.Create(book)
	require.Nil(t, err)
	require.Equal(t, book.Title, api.createBook.Title)
	require.Equal(t, book.Author, api.createBook.Author)
}

func TestDelete(t *testing.T) {
	api := &testBookApiService{deleteErr: nil}
	service := NewBookService(api)
	id := 1
	err := service.Delete(id)
	require.Nil(t, err)
	require.Equal(t, id, int(api.deleteId))
}

func TestGet(t *testing.T) {
	id := 1
	api := &testBookApiService{getBook: openapitools.Book{Id: int32(id), Title: "abc", Author: "xyz"}, getErr: nil}
	service := NewBookService(api)
	book, err := service.Get(id)
	require.Nil(t, err)
	require.Equal(t, int(api.getBook.Id), book.Id)
	require.Equal(t, api.getBook.Title, book.Title)
	require.Equal(t, api.getBook.Author, book.Author)
}

func TestList(t *testing.T) {
	dto := []openapitools.Book{
		{Id: 1, Title: "abc", Author: "def"},
		{Id: 2, Title: "ghi", Author: "jkl"},
	}
	api := &testBookApiService{listBooks: dto, listErr: nil}
	service := NewBookService(api)
	books, err := service.List()
	require.Nil(t, err)
	for i, book := range books {
		require.Equal(t, int(dto[i].Id), book.Id)
		require.Equal(t, dto[i].Title, book.Title)
		require.Equal(t, dto[i].Author, book.Author)
	}
}

func TestUpdate(t *testing.T) {
	api := &testBookApiService{updateErr: nil}
	service := NewBookService(api)
	book := data.Book{Id: 1, Title: "abc", Author: "xyz"}
	err := service.Update(book.Id, book)
	require.Nil(t, err)
	require.Equal(t, book.Id, int(api.updateId))
	require.Equal(t, book.Title, api.updateBook.Title)
	require.Equal(t, book.Author, api.updateBook.Author)
}

package openapitools

import (
	"context"
	"github.com/csandiego/poc-openapitools/go/client/data"
	"github.com/csandiego/poc-openapitools/go/openapitools"
)

type BookService struct {
	api BookApi
}

func NewBookService(api BookApi) *BookService {
	return &BookService{api: api}
}

func (s *BookService) Create(book data.Book) error {
	_, err := s.api.Create(context.Background(), openapitools.Book{Id: int32(book.Id), Title: book.Title, Author: book.Author})
	return err
}

func (s *BookService) Delete(id int) error {
	_, err := s.api.Delete(context.Background(), int32(id))
	return err
}

func (s *BookService) Get(id int) (data.Book, error) {
	dto, _, err := s.api.Get(context.Background(), int32(id))
	book := data.Book{}
	if err == nil {
		book.Id = int(dto.Id)
		book.Title = dto.Title
		book.Author = dto.Author
	}
	return book, err
}

func (s *BookService) List() ([]data.Book, error) {
	dto, _, err := s.api.List(context.Background())
	books := []data.Book{}
	if err == nil {
		for _, book := range dto {
			books = append(books, data.Book{Id: int(book.Id), Title: book.Title, Author: book.Author})
		}
	}
	return books, err
}

func (s *BookService) Update(id int, book data.Book) error {
	_, err := s.api.Update(context.Background(), int32(id), openapitools.Book{Id: int32(book.Id), Title: book.Title, Author: book.Author})
	return err
}

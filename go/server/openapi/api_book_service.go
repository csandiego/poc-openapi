/*
 * Books API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/csandiego/poc-openapitools/go/server/service"
)

// BookApiService is a service that implents the logic for the BookApiServicer
// This service should implement the business logic for every endpoint for the BookApi API. 
// Include any external packages or services that will be required by this service.
type BookApiService struct {
	service service.BookService
}

// NewBookApiService creates a default api service
func NewBookApiService(service service.BookService) BookApiServicer {
	return &BookApiService{service: service}
}

// Create - 
func (s *BookApiService) Create(book Book) (interface{}, error) {
	b := service.Book{Title: book.Title, Author: book.Author}
	return nil, s.service.Create(b)
}

// Delete - 
func (s *BookApiService) Delete(id int32) (interface{}, error) {
	return nil, s.service.Delete(int(id))
}

// Get - 
func (s *BookApiService) Get(id int32) (interface{}, error) {
	book, err := s.service.Get(int(id))
	if err != nil {
		return nil, err
	}
	return Book{Id: int32(book.Id), Title: book.Title, Author: book.Author}, nil
}

// List - 
func (s *BookApiService) List() (interface{}, error) {
	books, err := s.service.List()
	if err != nil {
		return nil, err
	}
	output := []Book{}
	for _, b := range books {
		output = append(output, Book{Id: int32(b.Id), Title: b.Title, Author: b.Author})
	}
	return output, nil
}

// Update - 
func (s *BookApiService) Update(id int32, book Book) (interface{}, error) {
	b := service.Book{Title: book.Title, Author: book.Author}
	return nil, s.service.Update(int(id), b)
}
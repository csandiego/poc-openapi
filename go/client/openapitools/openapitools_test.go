package openapitools

import (
	"context"
	"github.com/csandiego/poc-openapitools/go/openapitools"
	"net/http"
)

type testBookApiService struct {
	createBook openapitools.Book
	createErr  error
	deleteId   int32
	deleteErr  error
	getId      int32
	getBook    openapitools.Book
	getErr     error
	listBooks  []openapitools.Book
	listErr    error
	updateId   int32
	updateBook openapitools.Book
	updateErr  error
}

func (api *testBookApiService) Create(ctx context.Context, book openapitools.Book) (*http.Response, error) {
	api.createBook = book
	return nil, api.createErr
}

func (api *testBookApiService) Delete(ctx context.Context, id int32) (*http.Response, error) {
	api.deleteId = id
	return nil, api.deleteErr
}

func (api *testBookApiService) Get(ctx context.Context, id int32) (openapitools.Book, *http.Response, error) {
	api.getId = id
	return api.getBook, nil, api.getErr
}

func (api *testBookApiService) List(ctx context.Context) ([]openapitools.Book, *http.Response, error) {
	return api.listBooks, nil, api.listErr
}

func (api *testBookApiService) Update(ctx context.Context, id int32, book openapitools.Book) (*http.Response, error) {
	api.updateId = id
	api.updateBook = book
	return nil, api.updateErr
}

package openapitools

import (
	"context"
	"github.com/csandiego/poc-openapitools/go/openapitools"
	"net/http"
)

type BookApi interface {
	Create(context.Context, openapitools.Book) (*http.Response, error)
	Delete(context.Context, int32) (*http.Response, error)
	Get(context.Context, int32) (openapitools.Book, *http.Response, error)
	List(context.Context) ([]openapitools.Book, *http.Response, error)
	Update(context.Context, int32, openapitools.Book) (*http.Response, error)
}

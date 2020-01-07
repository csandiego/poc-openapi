package service

import (
	"errors"
	"log"
	"sync"
)

var ErrNotFound = errors.New("Not found")

type MockBookService struct {
	bookId int
	books  map[int]Book
	mutex  sync.Mutex
}

func NewMockBookService() *MockBookService {
	return &MockBookService{bookId: 0, books: map[int]Book{}}
}

func (service *MockBookService) Create(book Book) error {
	service.mutex.Lock()
	service.bookId += 1
	book.Id = service.bookId
	service.books[service.bookId] = book
	log.Println(service.books)
	service.mutex.Unlock()
	return nil
}

func (service *MockBookService) Delete(id int) error {
	service.mutex.Lock()
	delete(service.books, id)
	log.Println(service.books)
	service.mutex.Unlock()
	return nil
}

func (service *MockBookService) Get(id int) (Book, error) {
	var book Book
	var exists bool
	service.mutex.Lock()
	book, exists = service.books[id]
	service.mutex.Unlock()
	if !exists {
		return book, ErrNotFound
	}
	return book, nil
}

func (service *MockBookService) List() ([]Book, error) {
	books := []Book{}
	service.mutex.Lock()
	for _, v := range service.books {
		books = append(books, v)
	}
	service.mutex.Unlock()
	return books, nil
}

func (service *MockBookService) Update(id int, book Book) error {
	var err error = nil
	service.mutex.Lock()
	if _, exists := service.books[id]; exists {
		book.Id = id
		service.books[id] = book
	} else {
		err = ErrNotFound
	}
	log.Println(service.books)
	service.mutex.Unlock()
	return err
}

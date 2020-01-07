package service

type Book struct {
	Id     int
	Title  string
	Author string
}

type BookService interface {
	Create(Book) error
	Delete(int) error
	Get(int) (Book, error)
	List() ([]Book, error)
	Update(int, Book) error
}

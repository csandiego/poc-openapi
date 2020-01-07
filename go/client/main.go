package main

import (
	"github.com/csandiego/poc-openapitools/go/client/data"
	"github.com/csandiego/poc-openapitools/go/client/openapitools"
	"github.com/csandiego/poc-openapitools/go/client/service"
	oat "github.com/csandiego/poc-openapitools/go/openapitools"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Insufficient arguments.")
	}
	var service service.BookService = openapitools.NewBookService(oat.NewAPIClient(oat.NewConfiguration()).BookApi)
	switch os.Args[1] {
	case "list":
		books, err := service.List()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(books)
	case "create":
		if len(os.Args) != 4 {
			log.Fatal("Wrong number of arguments for create.")
		}
		if err := service.Create(data.Book{Title: os.Args[2], Author: os.Args[3]}); err != nil {
			log.Fatal(err)
		}
	case "get":
		if len(os.Args) != 3 {
			log.Fatal("Wrong number of arguments for get.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		book, err := service.Get(int(id))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(book)
	case "update":
		if len(os.Args) != 5 {
			log.Fatal("Wrong number of arguments for update.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if err = service.Update(int(id), data.Book{Title: os.Args[3], Author: os.Args[4]}); err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(os.Args) != 3 {
			log.Fatal("Wrong number of arguments for delete.")
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		if err = service.Delete(int(id)); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unsupported operation")
	}
}

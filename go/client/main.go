package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"github.com/csandiego/poc-openapitools/go/openapitools"
)

func main() {
	client := openapitools.NewAPIClient(openapitools.NewConfiguration())
	switch os.Args[1] {
	case "list":
		books, response, err := client.BookApi.List(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(books)
		log.Println(response.Status)
	case "create":
		response, err := client.BookApi.Create(context.Background(), openapitools.Book{Title: os.Args[2], Author: os.Args[3]})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response.Status)
	case "get":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		book, response, err := client.BookApi.Get(context.Background(), int32(id))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(book)
		log.Println(response.Status)
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		response, err := client.BookApi.Update(context.Background(), int32(id), openapitools.Book{Title: os.Args[3], Author: os.Args[4]})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response.Status)
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		response, err := client.BookApi.Delete(context.Background(), int32(id))
		if err != nil {
			log.Fatal(err)
		}
		log.Println(response.Status)
	default:
		log.Fatal("Unsupported operation")
	}
}

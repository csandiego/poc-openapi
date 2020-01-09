package com.github.csandiego.pocopenapitools.client.openapitools

import com.github.csandiego.pocopenapitools.client.data.Book
import com.github.csandiego.pocopenapitools.client.service.BookService
import com.github.csandiego.pocopenapitools.openapitools.models.Book as DTO

class OpenApiToolsBookService(private val api: OpenApiToolsBookApi) : BookService {

    override fun list(): List<Book> {
	return api.list().map { Book(it.id, it.title, it.author) }
    }

    override fun create(book: Book) {
	api.create(DTO(book.title, book.author, book.id))
    }

    override fun get(id: Int): Book {
	return api.get(id).run { Book(id, title, author) }
    }

    override fun update(id: Int, book: Book) {
	api.update(id, DTO(book.title, book.author, book.id))
    }

    override fun delete(id: Int) {
	api.delete(id)
    }
}

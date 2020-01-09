package com.github.csandiego.pocopenapitools.client.openapitools

import com.github.csandiego.pocopenapitools.client.data.Book
import com.github.csandiego.pocopenapitools.openapitools.models.Book as DTO
import org.junit.*
import org.junit.Assert.*

class OpenApiToolsBookServiceTest {

    private lateinit var api: TestBookApi
    private lateinit var service: OpenApiToolsBookService

    @Before
    fun setUp() {
	api = TestBookApi()
	service = OpenApiToolsBookService(api)
    }

    @Test
    fun list() {
	api.listBooks = arrayOf(DTO("abc", "xyz", 1))
	val books = service.list()
	assertEquals(api.listBooks.size, books.size)
	assertEquals(api.listBooks[0].id, books[0].id)
	assertEquals(api.listBooks[0].title, books[0].title)
	assertEquals(api.listBooks[0].author, books[0].author)
    }

    @Test
    fun create() {
	val book = Book(null, "abc", "xyz")
	service.create(book)
	assertEquals(book.title, api.createBook.title)
	assertEquals(book.author, api.createBook.author)
    }

    @Test
    fun get() {
	val id = 1
	api.getBook = DTO("abc", "xyz", id)
	val book = service.get(id)
	assertEquals(id, api.getId)
	assertEquals(api.getBook.id, book.id)
	assertEquals(api.getBook.title, book.title)
	assertEquals(api.getBook.author, book.author)
    }

    @Test
    fun update() {
	val id = 1
	val book = Book(null, "abc", "xyz")
	service.update(id, book)
	assertEquals(id, api.updateId)
	assertEquals(book.title, api.updateBook.title)
	assertEquals(book.author, api.updateBook.author)
    }

    @Test
    fun delete() {
	val id = 1
	service.delete(id)
	assertEquals(id, api.deleteId)
    }

    private class TestBookApi: OpenApiToolsBookApi {
	lateinit var listBooks: Array<DTO>
	lateinit var createBook: DTO
	var getId = 0
	lateinit var getBook: DTO
	var updateId = 0
	lateinit var updateBook: DTO
	var deleteId = 0

	override fun list() = listBooks

	override fun create(book: DTO) {
	    createBook = book
	}

	override fun get(id: Int): DTO {
	    getId = id
	    return getBook
	}

	override fun update(id: Int, book: DTO) {
	    updateId = id
	    updateBook = book
	}

	override fun delete(id: Int) {
	    deleteId = id
	}
    }
}

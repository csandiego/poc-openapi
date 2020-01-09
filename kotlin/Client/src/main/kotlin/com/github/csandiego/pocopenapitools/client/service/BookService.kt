package com.github.csandiego.pocopenapitools.client.service

import com.github.csandiego.pocopenapitools.client.data.Book

interface BookService {
    fun list(): List<Book>
    fun create(book: Book)
    fun get(id: Int): Book
    fun update(id: Int, book: Book)
    fun delete(id: Int)
}

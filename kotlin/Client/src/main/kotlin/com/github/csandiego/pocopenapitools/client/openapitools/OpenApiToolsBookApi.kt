package com.github.csandiego.pocopenapitools.client.openapitools

import com.github.csandiego.pocopenapitools.openapitools.models.Book

interface OpenApiToolsBookApi {

    fun list(): Array<Book>

    fun create(book: Book)

    fun get(id: Int): Book

    fun update(id: Int, book: Book)

    fun delete(id: Int)
    
}

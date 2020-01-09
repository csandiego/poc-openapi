package com.github.csandiego.pocopenapitools.client.openapitools

import com.github.csandiego.pocopenapitools.openapitools.apis.BookApi
import com.github.csandiego.pocopenapitools.openapitools.models.Book

class OpenApiToolsBookApiWrapper(private val api: BookApi) : OpenApiToolsBookApi {

    override fun list() = api.list()

    override fun create(book: Book) = api.create(book)

    override fun get(id: Int) = api.get(id)

    override fun update(id: Int, book: Book) = api.update(id, book)

    override fun delete(id: Int) = api.delete(id)
    
}

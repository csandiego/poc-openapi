/*
 * This Kotlin source file was generated by the Gradle 'init' task.
 */
package com.github.csandiego.pocopenapitools.client

import com.github.csandiego.pocopenapitools.client.data.Book
import com.github.csandiego.pocopenapitools.client.openapitools.OpenApiToolsBookApiWrapper
import com.github.csandiego.pocopenapitools.client.openapitools.OpenApiToolsBookService
import com.github.csandiego.pocopenapitools.openapitools.apis.BookApi
import kotlin.system.exitProcess

fun main(args: Array<String>) {
    if (args.size < 1) {
	println("Insufficient arguments.")
	exitProcess(1)
    }
    val service = OpenApiToolsBookService(OpenApiToolsBookApiWrapper(BookApi()))
    try {
	when (args[0]) {
            "list" -> println("[" + service.list().joinToString(",") + "]")
            "create" -> {
		if (args.size != 3) {
		    println("Wrong number of arguments for create.")
		    exitProcess(1)
		}
		service.create(Book(null, args[1], args[2]))
	    }
            "get" -> {
		if (args.size != 2) {
		    println("Wrong number of arguments for get.")
		    exitProcess(1)
		}
		println(service.get(args[1].toInt()))
	    }
            "update" -> {
		if (args.size != 4) {
		    println("Wrong number of arguments for update.")
		    exitProcess(1)
		}
		service.update(args[1].toInt(), Book(null, args[2], args[3]))
	    }
            "delete" -> {
		if (args.size != 2) {
		    println("Wrong number of arguments for delete.")
		    exitProcess(1)
		}
		service.delete(args[1].toInt())
	    }
            else -> println("Unsupported operation")

	}
    } catch (e: Exception) {
	println(e)
    }
}

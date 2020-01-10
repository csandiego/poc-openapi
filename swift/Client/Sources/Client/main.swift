import Foundation
import OpenApiTools
import Lib

if CommandLine.argc < 2 {
    print("Insufficient arguments")
    exit(EXIT_FAILURE)
}

OpenApiToolsAPI.apiResponseQueue = .global(qos: .`default`)

let service = OpenApiToolsBookService(api: BookAPI.self)

switch CommandLine.arguments[1] {
case "list":
    let group = DispatchGroup()
    group.enter()
    service.list { books, error in
        if let books = books {
            print(books)
        } else {
            print(error!)
        }
        group.leave()
    }
    group.wait()
case "create":
    if CommandLine.argc != 4 {
        print("Wrong number of arguments for create.")
        exit(EXIT_FAILURE)
    }
    let group = DispatchGroup()
    group.enter()
    service.create(book: Book(id: nil, title: CommandLine.arguments[2], author: CommandLine.arguments[3])) { _, error in
        if let error = error {
            print(error)
        }
        group.leave()
    }
    group.wait()
case "get":
    if CommandLine.argc != 3 {
        print("Wrong number of arguments for get.")
        exit(EXIT_FAILURE)
    }
    let id = Int(CommandLine.arguments[2])!
    let group = DispatchGroup()
    group.enter()
    service.callGet(id: id) { book, error in
        if let book = book {
            print(book)
        } else {
            print(error!)
        }
        group.leave()
    }
    group.wait()
case "update":
    if CommandLine.argc != 5 {
        print("Wrong number of arguments for update.")
        exit(EXIT_FAILURE)
    }
    let id = Int(CommandLine.arguments[2])!
    let group = DispatchGroup()
    group.enter()
    service.update(id: id, book: Book(id: nil, title: CommandLine.arguments[3], author: CommandLine.arguments[4])) { _, error in
        if let error = error {
            print(error)
        }
        group.leave()
    }
    group.wait()
case "delete":
    if CommandLine.argc != 3 {
        print("Wrong number of arguments for delete.")
        exit(EXIT_FAILURE)
    }
    let id = Int(CommandLine.arguments[2])!
    let group = DispatchGroup()
    group.enter()
    service.delete(id: id) { _, error in
        if let error = error {
            print(error)
        }
        group.leave()
    }
    group.wait()
default:
    print("Unsupported operation")
}

import Foundation
import OpenApiTools
import Lib

OpenApiToolsAPI.apiResponseQueue = .global(qos: .`default`)

let service = OpenApiToolsBookService(api: BookAPI.self)

switch CommandLine.arguments[1] {
case "list":
    let group = DispatchGroup()
    group.enter()
    service.list { books, error in
        if let error = error {
            print(error)
        }
        if let books = books {
            print(books)
        }
        group.leave()
    }
    group.wait()
case "create":
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
    guard let id = Int(CommandLine.arguments[2]) else {
        break
    }
    let group = DispatchGroup()
    group.enter()
    service.callGet(id: id) { book, error in
        if let error = error {
            print(error)
        } else {
            print(book!)
        }
        group.leave()
    }
    group.wait()
case "update":
    guard let id = Int(CommandLine.arguments[2]) else {
        break
    }
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
    guard let id = Int(CommandLine.arguments[2]) else {
        break
    }
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

import XCTest
import Lib
import OpenApiTools

class OpenApiToolsBookServiceTests: XCTestCase {

    var state: State!
    var service: OpenApiToolsBookService!

    override func setUp() {
        state = State()
        TestBookAPI.state = state
        service = OpenApiToolsBookService(api: TestBookAPI.self)
    }

    func testList() {
        state.listBooks = [OpenApiTools.Book(id: 1, title: "abc", author: "xyz")]
        let expectation = self.expectation(description: "Callback")
        service.list { books, error in
            if let books = books {
                XCTAssertEqual(self.state.listBooks[0].id, books[0].id)
                XCTAssertEqual(self.state.listBooks[0].title, books[0].title)
                XCTAssertEqual(self.state.listBooks[0].author, books[0].author)
                expectation.fulfill()
            } else {
                XCTFail()
            }
        }
        wait(for: [expectation], timeout: 1.0)
    }

    func testCreate() {
        let book = Lib.Book(id: 1, title: "abc", author: "xyz")
        let expectation = self.expectation(description: "Callback")
        service.create(book: book) { _, error in
            XCTAssertNil(error)
            XCTAssertEqual(book.id, self.state.createBook.id)
            XCTAssertEqual(book.title, self.state.createBook.title)
            XCTAssertEqual(book.author, self.state.createBook.author)
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }

    func testGet() {
        let id = 1
        state.getBook = OpenApiTools.Book(id: id, title: "abc", author: "xyz")
        let expectation = self.expectation(description: "Callback")
        service.callGet(id: id) { book, error in
            XCTAssertNil(error)
            XCTAssertEqual(id, self.state.getId)
            guard let book = book else {
                XCTFail()
                return
            }
            XCTAssertEqual(self.state.getBook.id, book.id)
            XCTAssertEqual(self.state.getBook.title, book.title)
            XCTAssertEqual(self.state.getBook.author, book.author)
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }

    func testUpdate() {
        let id = 1
        let book = Lib.Book(id: nil, title: "abc", author: "xyz")
        let expectation = self.expectation(description: "Callback")
        service.update(id: id, book: book) { _, error in
            XCTAssertNil(error)
            XCTAssertEqual(id, self.state.updateId)
            XCTAssertEqual(book.title, self.state.updateBook.title)
            XCTAssertEqual(book.author, self.state.updateBook.author)
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }

    func testDelete() {
        let id = 1
        let expectation = self.expectation(description: "Callback")
        service.delete(id: id) { _, error in
            XCTAssertNil(error)
            XCTAssertEqual(id, self.state.deleteId)
            expectation.fulfill()
        }
        wait(for: [expectation], timeout: 1.0)
    }

    class State {
        var getId = 0
        var getBook: OpenApiTools.Book!
        var createBook: OpenApiTools.Book!
        var deleteId = 0
        var listBooks: [OpenApiTools.Book]!
        var updateId = 0
        var updateBook: OpenApiTools.Book!
    }

    class TestBookAPI: OpenApiToolsBookAPI {

        static var state: State!

        static func callGet(id: Int, completion: @escaping ((_ data: OpenApiTools.Book?, _ error: Error?) -> Void)) {
            state.getId = id
            completion(state.getBook, nil)
        }

        static func create(book: OpenApiTools.Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
            state.createBook = book
            completion(nil, nil)
        }

        static func delete(id: Int, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
            state.deleteId = id
            completion(nil, nil)
        }

        static func list(completion: @escaping ((_ data: [OpenApiTools.Book]?, _ error: Error?) -> Void)) {
            completion(state.listBooks, nil)
        }

        static func update(id: Int, book: OpenApiTools.Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
            state.updateId = id
            state.updateBook = book
            completion(nil, nil)
        }
    
    }
    
}

import OpenApiTools

public class OpenApiToolsBookService: BookService {

    let api: OpenApiToolsBookAPI.Type

    public init(api: OpenApiToolsBookAPI.Type) {
        self.api = api
    }

    public func list(completion: @escaping ((_ books: [Book]?, _ error: Error?) -> Void)) {
        api.list { dto, error in
            completion(dto?.map { Book(id: $0.id, title: $0.title, author: $0.author) }, error)
        }
    }

    public func create(book: Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
        api.create(book: OpenApiTools.Book(id: book.id, title: book.title, author: book.author)) { data, error in
            completion(data, error)
        }
    }

    public func callGet(id: Int, completion: @escaping ((_ book: Book?, _ error: Error?) -> Void)) {
        api.callGet(id: id) { dto, error in
            var book: Book? = nil
            if let dto = dto {
                book = Book(id: dto.id, title: dto.title, author: dto.author)
            }
            completion(book, error)
        }
    }

    public func update(id: Int, book: Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
        api.update(id: id, book: OpenApiTools.Book(id: book.id, title: book.title, author: book.author)) { data, error in
            completion(data, error)
        }
    }

    public func delete(id: Int, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void)) {
        api.delete(id: id) { data, error in
            completion(data, error)
        }
    }
    
}

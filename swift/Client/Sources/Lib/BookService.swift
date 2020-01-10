public protocol BookService {

    func list(completion: @escaping ((_ books: [Book]?, _ error: Error?) -> Void))

    func create(book: Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))

    func callGet(id: Int, completion: @escaping ((_ book: Book?, _ error: Error?) -> Void))

    func update(id: Int, book: Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))

    func delete(id: Int, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))
    
}

import OpenApiTools

public protocol OpenApiToolsBookAPI {

    static func callGet(id: Int, completion: @escaping ((_ data: OpenApiTools.Book?, _ error: Error?) -> Void))

    static func create(book: OpenApiTools.Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))

    static func delete(id: Int, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))

    static func list(completion: @escaping ((_ data: [OpenApiTools.Book]?, _ error: Error?) -> Void))

    static func update(id: Int, book: OpenApiTools.Book, completion: @escaping ((_ data: Void?, _ error: Error?) -> Void))

}

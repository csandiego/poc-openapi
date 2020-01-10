public struct Book {
    
    public let id: Int?
    public let title: String
    public let author: String

    public init(id: Int?, title: String, author: String) {
        self.id = id
        self.title = title
        self.author = author
    }
    
}

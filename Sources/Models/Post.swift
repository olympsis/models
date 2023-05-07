//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/4/23.
//

import SwiftBSON
import Foundation

public struct Post: Codable, Identifiable {
    public var id: BSONObjectID?            // 6155a18d06eeab8e7559fc35 (mongodb)
    public var clubId: BSONObjectID         // 6155a18d06eeab8e7559fc35 (mongodb)
    public var poster: String               // 6abfe36c-5eea-4bb9-af6a-2aeee37af629
    public var title: String                // title
    public var body: String                 // body of post
    public var event: BSONObjectID?         // 6155a18d06eeab8e7559fc35 (mongodb)
    public var images: [String]?            // array of images
    public var likes: [Like]?               // likes
    public var comments: [Comment]?         // comments
    public var createdAt: Int64?            // 1682785470
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case clubId
        case poster
        case title
        case body
        case event
        case images
        case likes
        case comments
        case createdAt
    }
}

public struct Comment: Codable {
    public var id: BSONObjectID?            // 6155a18d06eeab8e7559fc35 (mongodb)
    public var uuid: String                 // 6abfe36c-5eea-4bb9-af6a-2aeee37af629
    public var text: String
    public var createdAt: Int64?
    
    public init(uuid: String, text: String) {
        self.uuid = uuid
        self.text = text
        self.createdAt = 0
    }
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case uuid
        case text
        case createdAt
    }
}


public struct PostsResponse: Codable {
    public var posts: [Post]
    public var totalPosts: Int
    
    public init(_ posts: [Post], totalPosts: Int) {
        self.posts = posts
        self.totalPosts = totalPosts
    }
}

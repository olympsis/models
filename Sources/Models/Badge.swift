//
//  Badge.swift
//  
//
//  Created by Joel Joseph on 5/13/23.
//

import SwiftBSON
import Foundation

public struct Badge: Codable {
    public let id: BSONObjectID?
    public let name: String
    public let title: String
    public let imageURL: String
    public let eventId: BSONObjectID?
    public let description: String
    public let achievedAt: Int64
    
    public init(id: BSONObjectID? = nil, name: String, title: String, imageURL: String, eventId: BSONObjectID? = nil, description: String, achievedAt: Int64) {
        self.id = id
        self.name = name
        self.title = title
        self.imageURL = imageURL
        self.eventId = eventId
        self.description = description
        self.achievedAt = achievedAt
    }
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case name
        case title
        case imageURL = "image_url"
        case eventId
        case description
        case achievedAt = "achieved_at"
    }
}

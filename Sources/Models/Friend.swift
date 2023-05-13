//
//  Friend.swift
//  
//
//  Created by Joel Joseph on 5/13/23.
//

import SwiftBSON
import Foundation

public struct Friend: Codable {
    public let uuid: String
    public let createdAt: Int64
    
    init(uuid: String, createdAt: Int64) {
        self.uuid = uuid
        self.createdAt = createdAt
    }
    
    enum CodingKeys: String, CodingKey {
        case uuid
        case createdAt = "created_at"
    }
}

public struct FriendRequest: Codable {
    public let id: BSONObjectID?
    public let requestor: String
    public let requestee: String
    public let status: String
    public let createdAt: Int64?
    
    public init(id: BSONObjectID? = nil, requestor: String, requestee: String, status: String, createdAt: Int64?=nil) {
        self.id = id
        self.requestor = requestor
        self.requestee = requestee
        self.status = status
        self.createdAt = createdAt
    }
    
    enum CodingKeys: String, CodingKey {
        case id
        case requestor
        case requestee
        case status
        case createdAt = "created_at"
    }
}

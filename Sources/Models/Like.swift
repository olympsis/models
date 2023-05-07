//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/4/23.
//

import SwiftBSON
import Foundation

/// Like object to keep track of how many people like the event/post
///
/// - user unique id
public struct Like: Codable {
    public var id: BSONObjectID?
    public var uuid: String                 // 6abfe36c-5eea-4bb9-af6a-2aeee37af629
    public var createdAt: Int64             // 1682785470
    
    public init(uuid: String, createdAt: Int64) {
        self.uuid = uuid
        self.createdAt = createdAt
    }
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case uuid
        case createdAt
    }
}

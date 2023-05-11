//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/2/23.
//
import SwiftBSON
import Foundation

/// User metadata
/// Does not contain user identifiable data
/// All data is related to the uuid which is related to auth user data which is identifiable
public struct User: Codable {
    public var uuid: String
    public var userName: String
    public var bio: String?
    public var imageURL: String?
    public var visibility: String
    public var clubs: [String]?
    public var sports: [String]?
    public var deviceToken: String?

    public init(uuid: String, userName: String, bio: String? = nil, imageURL: String? = nil, visibility: String, clubs: [String]? = nil, sports: [String]? = nil, deviceToken: String? = nil) {
        self.uuid = uuid
        self.userName = userName
        self.bio = bio
        self.imageURL = imageURL
        self.visibility = visibility
        self.clubs = clubs
        self.sports = sports
        self.deviceToken = deviceToken
    }

    enum CodingKeys: String, CodingKey {
        case uuid
        case userName = "username"
        case bio
        case imageURL = "image_url"
        case visibility
        case clubs
        case sports
        case deviceToken = "device_token"
    }
}

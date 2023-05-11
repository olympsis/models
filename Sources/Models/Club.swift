//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/4/23.
//

import SwiftBSON
import Foundation

public struct Club: Codable, Identifiable {
    public var id: BSONObjectID?
    public var name: String
    public var description: String
    public var sport: String
    public var city: String
    public var state: String
    public var country: String
    public var imageURL: String
    public var visibility: String
    public var members: [Member]?
    public var createdAt: Int64?
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case name
        case description
        case sport
        case city
        case state
        case country
        case imageURL = "image_url"
        case visibility
        case members
        case createdAt = "created_at"
    }
}

public struct Member: Codable, Identifiable {
    public var id: BSONObjectID?
    public var uuid: String
    public var role: String
    public var joinedAt: Int64?
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case uuid
        case role
        case joinedAt = "joined_at"
    }
}

public struct ClubResponse: Codable {
    public var token: String?
    public var club: String
}

public struct ClubsResponse: Codable {
    public var clubs: [Club]
    public var totalClubs: Int
    
    public init(_ clubs: [Club], totalClubs: Int) {
        self.clubs = clubs
        self.totalClubs = totalClubs
    }
}

public struct ClubApplication: Codable, Identifiable {
    public var id: BSONObjectID?
    public var clubId: BSONObjectID
    public var uuid: String
    public var status: String
    public var createdAt: Int64?
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case clubId = "club_id"
        case uuid
        case status
        case createdAt = "created_at"
    }
}

public struct ClubInvitation: Codable, Identifiable {
    public var id: BSONObjectID?
    public var clubId: BSONObjectID
    public var uuid: String
    public var status: String
    public var createdAt: Int64?
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case clubId = "club_id"
        case uuid
        case status
        case createdAt = "created_at"
    }
}

public struct ApplicationsRequest: Codable {
    public var type: String // club-user
    public var id: String // uuid-clubId
}

public struct ClubApplicationsReponse: Codable {
    public var applications: [ClubApplication]
    public var totalApplications: Int
    
    public init(applications: [ClubApplication], totalApplications: Int) {
        self.applications = applications
        self.totalApplications = totalApplications
    }
}

public struct ClubInvitationsReponse: Codable {
    public var invitations: [ClubInvitation]
    public var totalInvitations: Int
    
    public init(invitations: [ClubInvitation], totalInvitations: Int) {
        self.invitations = invitations
        self.totalInvitations = totalInvitations
    }
}

public struct RoleRequest: Codable {
    public var role: String
}

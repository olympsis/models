//
//  Event.swift
//  
//
//  Created by Joel Joseph on 4/29/23.
//

import SwiftBSON
import Foundation

/// Data object to represent event
///
/// Contains all data pertaining to events such as:
/// - name
/// - field
/// - club
/// etc..
public struct Event: Codable, Identifiable {
    public var id: BSONObjectID?            // 6155a18d06eeab8e7559fc35 (mongodb)
    public var poster: String               // 6abfe36c-5eea-4bb9-af6a-2aeee37af629
    public var club: BSONObjectID           // 6155a18d06eeab8e7559fc35 (mongodb)
    public var field: BSONObjectID          // 6155a18d06eeab8e7559fc35 (mongodb)
    public var image: String                // soccer-0
    public var title: String                // Pick up
    public var body: String                 // Let's go play guys
    public var sport: String                // soccer
    public var level: String                // intermediate
    public var status: String               // pending
    public var startTime: Int64             // 1682785470
    public var aStartTime: Int64            // 1682785470
    public var stopTime: Int64              // 1682785470
    public var aStopTime: Int64             // 1682785470
    public var participants:[Participant]   // {}
    public var maxParticipants: Int         // 15
    public var likes: [Like]                // {}
    public var visibility: String           // public/private/club
    public var createdAt: Int64             // 1682785470
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case poster
        case club
        case field
        case image
        case title
        case body
        case sport
        case level
        case status
        case startTime
        case aStartTime
        case stopTime
        case aStopTime
        case participants
        case maxParticipants
        case likes
        case visibility
        case createdAt
    }
}

/// Participant object to keep track of who responded to the event
///
/// - user unique id
/// - response
public struct Participant: Codable, Identifiable {
    public var id: BSONObjectID             // 6155a18d06eeab8e7559fc35
    public var uuid: String                 // 6abfe36c-5eea-4bb9-af6a-2aeee37af629
    public var response: String             // yes/no/maybe
    public var createdAt: Int64             // 1682785470
    
    public init(id: BSONObjectID, uuid: String, response: String, createdAt: Int64) {
        self.id = id
        self.uuid = uuid
        self.response = response
        self.createdAt = createdAt
    }
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case uuid
        case response
        case createdAt
    }
}



/// Response object for get events request
///
/// - Contains an integer for the number of events
/// - Contains an array containing events objects
public struct EventsResponse: Codable {
    public var totalEvents: Int // total events num
    public var events: [Event]  // events array
    
    public init(_ events: [Event], totalEvents: Int){
        self.totalEvents = totalEvents
        self.events = events
    }
}

/// Response object for create event request
///
/// - Contains an id of the created event document
public struct CreateEventResponse: Codable, Hashable {
    public var id: BSONObjectID // event/document id
    
    public init(_ id: BSONObjectID) {
        self.id = id
    }
}

/// Request object for participants
///
/// - Contains the uuid of the participant
public struct ParticpantRequest: Codable {
    public var uuid: String
    public var response: String?
}

public struct LikeRequest: Codable {
    public var uuid: String
}

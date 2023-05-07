//
//  Field.swift
//  
//
//  Created by Joel Joseph on 4/27/23.
//

import SwiftBSON
import Foundation

/// Data object to represent fields
///
/// Contains all data pertaining to fields such as:
/// - name
/// - location
/// - images
/// etc..
public struct Field: Codable, Hashable, Identifiable {
    public var id: BSONObjectID?            // 6155a18d06eeab8e7559fc35 (mongodb)
    public var name: String                 // Richard Building Fields
    public var owner: Ownership             // {"name":"BrighamYoung","type": "private"}
    public var description: String          // lorem ipsum....
    public var sports: [String]             // ["soccer", "basketball"]
    public var images: [String]             // /field-images/xxxxxx.jpg
    public var location: GeoJSON            // {"type":"point","cord":[00.00, 00.00]}
    public var city: String                 // Provo
    public var state: String                // UT
    public var country: String              // United States of America
    
    enum CodingKeys: String, CodingKey {
        case id = "_id"
        case name
        case owner
        case description
        case sports
        case images
        case location
        case city
        case state
        case country
    }
}

/// Geojson object for specifying coordinates only points
///
/// - Geojson type
/// - Coordinates array for location points
public struct GeoJSON: Codable, Hashable {
    public var type: String             // point
    public var coordinates: [Double]    // long/lat (2)
}

/// Object to define ownership of a field
///
/// - Field owner legal name
/// - Field owner ownership type
public struct Ownership: Codable, Hashable {
    public var name: String    // owner name
    public var type: String    // public-private
}

/// Response object for get fields request
///
/// - Contains an integer for the number of fields
/// - Contains an array containing field objects
public struct FieldsResponse: Codable, Hashable {
    public var totalFields: Int // total fields num
    public var fields: [Field]  // fields array
    
    public init(_ fields: [Field], totalFields: Int){
        self.totalFields = totalFields
        self.fields = fields
    }
}

/// Response object for create fields request
///
/// - Contains an id of the created field document
public struct CreateResponse: Codable, Hashable {
    public var id: BSONObjectID // field/document id
    
    public init(_ id: BSONObjectID) {
        self.id = id
    }
}

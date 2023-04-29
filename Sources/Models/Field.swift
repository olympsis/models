//
//  File.swift
//  
//
//  Created by Joel Joseph on 4/27/23.
//

import SwiftBSON
import Foundation

public struct Field: Codable, Hashable, Identifiable {
    public var id: BSONObjectID           // id       - 6155a18d06eeab8e7559fc35 (mongodb)
    public var name: String         // name     - Richard Building Fields
    public var owner: Ownership     // owner    - {"name":"BrighamYoung","type": "private"}
    public var description: String  // desc     - lorem ipsum....
    public var sports: [String]     // sports   - ["soccer", "basketball"]
    public var images: [String]     // images   - /field-images/xxxxxx.jpg
    public var location: GeoJSON    // loc      - {"type":"point","cord":[00.00, 00.00]}
    public var city: String         // city     - Provo
    public var state: String        // state    - UT
    public var country: String      // country  - United States of America
    
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

public struct GeoJSON: Codable, Hashable {
    var type: String    // point
    var coordinates: [Double] // long/lat (2)
}

public struct Ownership: Codable, Hashable {
    var name: String    // owner name
    var type: String    // public-private
}

public struct FieldsResponse: Codable, Hashable {
    var totalFields: Int
    var fields: [Field]
    
    public init(_ fields: [Field], totalFields: Int){
        self.totalFields = totalFields
        self.fields = fields
    }
}

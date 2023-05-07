//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/2/23.
//

import Foundation

public struct User: Codable {
    public let uuid: String
    public let userName: String
    public let bio: String?
    public let imageURL: String?
    public let visibility: String
    public let clubs: [String]?
    public let sports: [String]?
    public let deviceToken: String?
}


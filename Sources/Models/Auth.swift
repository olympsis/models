//
//  File.swift
//  
//
//  Created by Joel Joseph on 5/10/23.
//

import SwiftBSON
import Foundation

/// Auth user data
/// Contains user identifiable data
public struct AuthUser: Codable {
    public var uuid: String
    public var firstName: String
    public var lastName: String
    public var email: String
    public var token: String
    public var accessToken: String
    public var provider: String
    public var createdAt: Int64
    
    public init(uuid: String, firstName: String, lastName: String, email: String, token: String, accessToken: String, provider: String, createdAt: Int64) {
        self.uuid = uuid
        self.firstName = firstName
        self.lastName = lastName
        self.email = email
        self.token = token
        self.accessToken = accessToken
        self.provider = provider
        self.createdAt = createdAt
    }
    
    enum CodingKeys: String, CodingKey {
        case uuid
        case firstName = "first_name"
        case lastName = "last_name"
        case email
        case token
        case accessToken = "access_token"
        case provider
        case createdAt = "created_at"
    }
}

/// Auth request
/// Contains user identifiable data to talk to backend
public struct AuthRequest: Codable {
    public var firstName: String?
    public var lastName: String?
    public var email: String?
    public var code: String
    public var provider: String
    
    // Login Request
    public init(code: String, provider: String) {
        self.code = code
        self.provider = provider
    }
    
    // Signin request
    public init(firstName: String, lastName: String, email: String, code: String, provider: String) {
        self.firstName = firstName
        self.lastName = lastName
        self.email = email
        self.code = code
        self.provider = provider
    }
    
    enum CodingKeys: String, CodingKey {
        case firstName = "first_name"
        case lastName = "last_name"
        case email
        case code
        case provider
    }
}

/// Auth response
/// Contains user identifiable data to respond from auth request
public struct AuthResponse: Codable {
    public var uuid: String
    public var firstName: String
    public var lastName: String
    public var email: String
    public var token: String
    
    public init(uuid: String, firstName: String, lastName: String, email: String, token: String) {
        self.uuid = uuid
        self.firstName = firstName
        self.lastName = lastName
        self.email = email
        self.token = token
    }
    
    enum CodingKeys: String, CodingKey {
        case uuid
        case firstName = "first_name"
        case lastName = "last_name"
        case email
        case token
    }
}


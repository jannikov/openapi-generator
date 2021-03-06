//
// MapTest.swift
//
// Generated by openapi-generator
// https://openapi-generator.tech
//

import Foundation

internal struct MapTest: Codable {

    internal enum MapOfEnumString: String, Codable, CaseIterable {
        case upper = "UPPER"
        case lower = "lower"
    }
    internal var mapMapOfString: [String: [String: String]]?
    internal var mapOfEnumString: [String: String]?
    internal var directMap: [String: Bool]?
    internal var indirectMap: StringBooleanMap?

    internal init(mapMapOfString: [String: [String: String]]?, mapOfEnumString: [String: String]?, directMap: [String: Bool]?, indirectMap: StringBooleanMap?) {
        self.mapMapOfString = mapMapOfString
        self.mapOfEnumString = mapOfEnumString
        self.directMap = directMap
        self.indirectMap = indirectMap
    }

    internal enum CodingKeys: String, CodingKey, CaseIterable {
        case mapMapOfString = "map_map_of_string"
        case mapOfEnumString = "map_of_enum_string"
        case directMap = "direct_map"
        case indirectMap = "indirect_map"
    }

}

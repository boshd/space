//
//  APIService.swift
//  spaceios
//
//  Created by Kareem Arab on 2019-04-14.
//  Copyright © 2019 Kareem Arab. All rights reserved.
//

import Foundation
import Alamofire

typealias GETAPIResponse = (DataResponse<Data>) -> Void

struct APIService {
    
    static func getAPIResponse(_ url: String, username: String, password: String, completion: @escaping GETAPIResponse) {
        let parameters: [String: String] = [
            "username": username,
            "password": password
        ]
        
        Alamofire.request(url, method: .post, parameters: parameters)
            .responseJSON { response in
                print(response.request)  // original URL request
                print(response.response) // URL response
                print(response.data)     // server data
                print(response.result)   // result of response serialization
                
                if let JSON = response.result.value {
                    print("JSON: \(JSON)")
                }
        }
    }
    
}

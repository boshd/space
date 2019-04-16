//
//  ViewController.swift
//  spaceios
//
//  Created by Kareem Arab on 2019-04-14.
//  Copyright © 2019 Kareem Arab. All rights reserved.
//

import UIKit

class ViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        
        let url = "http://localhost:3000/api/user/login"
        APIService.getAPIResponse(url, username: "kareemd", password: "lecfej") { (s) in
            
        }
    }


}


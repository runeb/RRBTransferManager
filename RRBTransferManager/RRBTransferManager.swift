//
//  RRBTransferManager.swift
//  RRBTransferManager
//
//  Created by Rune Botten on 08.04.15.
//  Copyright (c) 2015 Rune Botten. All rights reserved.
//
import Foundation
import RRBTransferManager.TMCache

private let RRBTransferManagerIdentifier = "com.runerb.RRBTransferManager"
private let _RRBTransferManagerSingleton = RRBTransferManager()

public class RRBTransferManager: NSObject {
  internal let cache:TMCache

  public class var sharedInstance: RRBTransferManager {
    return _RRBTransferManagerSingleton
  }

  override init() {
    let cachePath = NSSearchPathForDirectoriesInDomains(.CachesDirectory, .UserDomainMask, true)[0] as NSString
    self.cache = TMCache(name: "\(RRBTransferManagerIdentifier).Cache", rootPath: cachePath)
  }
}

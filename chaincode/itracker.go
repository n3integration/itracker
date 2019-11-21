/*
 *  Copyright (C) 2019 n3integration
 *
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program. If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"log"
)

func init() {
	log.SetFlags(log.LUTC)
	log.SetPrefix("InventoryTracking")
}

// InventoryTrackingChaincode implementation
type InventoryTrackingChaincode struct {
	testMode bool
}

func (t *InventoryTrackingChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("initializing...")
	_, _ = stub.GetFunctionAndParameters()
	return shim.Success(nil)
}

func (t *InventoryTrackingChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("invoked")
	_, _ = stub.GetFunctionAndParameters()
	return shim.Error("invalid invoke function name")
}

func main() {
	itc := new(InventoryTrackingChaincode)
	itc.testMode = false
	err := shim.Start(itc)
	if err != nil {
		log.Printf("error starting InventoryTracking chaincode: %s", err)
	}
}

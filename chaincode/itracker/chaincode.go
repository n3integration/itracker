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
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	fnAdd          = "add"
	fnGet          = "get"
	fnGetHistory   = "history"
	fnQuery        = "query"
	fnTransfer     = "transfer"
	fnUpdateStatus = "update"
)

// InventoryTrackingChaincode implementation
type InventoryTrackingChaincode struct {
}

func (t *InventoryTrackingChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("initializing...")
	args := stub.GetStringArgs()
	if len(args) == 1 &&  args[0] == "testdata" {
		if err := loadTestData(stub); err != nil {
			return shim.Error(errors.Wrap(err, "failed to load test data").Error())
		}
	}

	return shim.Success(nil)
}

func (t *InventoryTrackingChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	log.Println(fmt.Sprintf("invoke %s...", fn))

	switch fn {
	case fnAdd:
		return add(stub, args...)
	case fnUpdateStatus:
		return updateStatus(stub, args...)
	case fnTransfer:
		return transfer(stub, args...)
	case fnGet:
		return get(stub, args...)
	case fnGetHistory:
		return getHistory(stub, args...)
	case fnQuery:
		return query(stub, args...)
	default:
		return newError(http.StatusBadRequest, "%q is not a supported operation", fn)
	}
}

func main() {
	itc := new(InventoryTrackingChaincode)
	err := shim.Start(itc)
	if err != nil {
		log.Printf("error starting chaincode: %s", err)
	}
}

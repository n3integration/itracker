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
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

func add(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 1 {
		return newError(http.StatusBadRequest, "item is required")
	}

	var item Item
	if err := json.NewDecoder(strings.NewReader(args[0])).Decode(&item); err != nil {
		return newError(http.StatusBadRequest, errors.Wrap(err, "malformed input").Error())
	}

	if err := item.Validate(); err != nil {
		return newError(http.StatusBadRequest, err.Error())
	}

	serializedItem, err := stub.GetState(item.Serial)
	if err != nil {
		return shim.Error(errors.Wrap(err, "failed to get item").Error())
	} else if serializedItem != nil {
		return newError(http.StatusConflict, "item with serial number %q already exists", item.Serial)
	}

	if cert, err := cid.GetX509Certificate(stub); err == nil {
		item.SubmittedBy = cert.Subject.CommonName
	}

	item.Status = Available
	serializedItem, _ = json.Marshal(item)
	if err := stub.PutState(item.Serial, serializedItem); err != nil {
		return shim.Error(errors.Wrap(err, "failed to save item").Error())
	}

	return shim.Success(serializedItem)
}

func updateStatus(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 1 {
		return newError(http.StatusBadRequest, "item serial number is required")
	}

	item, err := getState(stub, args[0])
	if item == nil {
		return err
	}

	if cert, err := cid.GetX509Certificate(stub); err == nil {
		item.SubmittedBy = cert.Subject.CommonName
	}

	item.Status = Unavailable
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(item)
	if err := stub.PutState(item.Serial, buffer.Bytes()); err != nil {
		return shim.Error(errors.Wrap(err, "failed to save item").Error())
	}

	return shim.Success(buffer.Bytes())
}

func transfer(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 2 {
		return newError(http.StatusBadRequest, "item serial number and facility are required")
	} else if args[1] == "" {
		return newError(http.StatusBadRequest, "item facility is required")
	}

	item, err := getState(stub, args[0])
	if item == nil {
		return err
	}

	if item.Status == Unavailable {
		return newError(http.StatusBadRequest, "item is no longer available and cannot be transferred")
	} else if item.Facility == args[1] {
		return shim.Success(nil)
	}

	if msp, err := cid.GetMSPID(stub); err != nil {
		return newError(http.StatusInternalServerError, "unable to validate request")
	} else if !strings.HasPrefix(msp, item.Facility) && !strings.HasPrefix(msp, args[1]) {
		return newError(http.StatusBadRequest, "item cannot be transferred to %q by %q", args[1], msp)
	}

	if cert, err := cid.GetX509Certificate(stub); err == nil {
		item.SubmittedBy = cert.Subject.CommonName
	}

	item.Facility = args[1]
	buffer := new(bytes.Buffer)
	_ = json.NewEncoder(buffer).Encode(item)
	if err := stub.PutState(item.Serial, buffer.Bytes()); err != nil {
		return shim.Error(errors.Wrap(err, "failed to save item").Error())
	}

	return shim.Success(buffer.Bytes())
}

func get(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 1 {
		return newError(http.StatusBadRequest, "item serial number is required")
	}

	serializedItem, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(errors.Wrap(err, "failed to get item").Error())
	} else if serializedItem == nil {
		return newError(http.StatusNotFound, "failed to find %s", args[0])
	}

	return shim.Success(serializedItem)
}

func getHistory(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	if len(args) != 1 {
		return newError(http.StatusBadRequest, "item serial number is required")
	}

	iterator, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(errors.Wrap(err, "failed to get item history").Error())
	}

	defer iterator.Close()

	first := true
	buffer := new(bytes.Buffer)
	fmt.Fprint(buffer, "[")
	for iterator.HasNext() {
		history, err := iterator.Next()
		if err != nil {
			return shim.Error(errors.Wrap(err, "failed to iterate over item history").Error())
		}

		if first {
			first = false
		} else {
			fmt.Fprint(buffer, ",")
		}

		fmt.Fprintf(buffer, `{"txId":%q,"value":%s,"timestamp":%q}`,
			history.TxId,
			toValue(history),
			time.Unix(history.Timestamp.Seconds, int64(history.Timestamp.Nanos)),
		)
	}

	fmt.Fprint(buffer, "]")
	return shim.Success(buffer.Bytes())
}

// by_category
// by_code
// by_facility
// by_manufacturer
// by_status
// limit
func query(stub shim.ChaincodeStubInterface, args ...string) peer.Response {
	query := `{"selector":{"_id":{"$gt":null}}}`
	if len(args) == 1 {
		query = args[0]
	}

	iterator, err := stub.GetQueryResult(query)
	if err != nil {
		return shim.Error(errors.Wrap(err, "failed to get item history").Error())
	}

	defer iterator.Close()

	first := true
	buffer := new(bytes.Buffer)
	fmt.Fprint(buffer, "[")
	for iterator.HasNext() {
		history, err := iterator.Next()
		if err != nil {
			return shim.Error(errors.Wrap(err, "failed to iterate over item history").Error())
		}

		if first {
			first = false
		} else {
			fmt.Fprint(buffer, ",")
		}

		fmt.Fprint(buffer, string(history.Value))
	}

	fmt.Fprint(buffer, "]")
	return shim.Success(buffer.Bytes())
}

func getState(stub shim.ChaincodeStubInterface, key string) (*Item, peer.Response) {
	serializedItem, err := stub.GetState(key)
	if err != nil {
		return nil, shim.Error(errors.Wrap(err, "failed to get item").Error())
	} else if serializedItem == nil {
		return nil, newError(http.StatusNotFound, "failed to find %s", key)
	}

	var item Item
	if err := json.NewDecoder(bytes.NewReader(serializedItem)).Decode(&item); err != nil {
		return nil, newError(http.StatusBadRequest, errors.Wrap(err, "malformed item state").Error())
	}
	return &item, shim.Success(nil)
}

type ChaincodeFunc func(shim.ChaincodeStubInterface, ...string) peer.Response

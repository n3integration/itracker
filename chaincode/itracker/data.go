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
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
	"github.com/pkg/errors"
)

func loadTestData(stub shim.ChaincodeStubInterface) error {
	facility := "Warehouse"
	submittedBy := "n3integration"
	if cert, err := cid.GetX509Certificate(stub); err == nil {
		submittedBy = cert.Subject.CommonName
	}

	items := []Item{
		{"", "SPOD", "POD300-ARB", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A00025", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A00029", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A20055", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A90007", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A20052", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A00027", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "VIAIR", "V/A20053", "Air Accessory Kit", facility, Available, submittedBy},
		{"", "Smittybilt", "S/B2781BAG", "Air Compressor Carry Bag", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92623", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92630", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92626", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92635", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92622", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92625", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92627", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92595", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92631", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "VIAIR", "V/A92621", "Air Compressor Filters", facility, Available, submittedBy},
		{"", "TeraFlex", "TER1184120", "Air Compressor Mounting Bracket", facility, Available, submittedBy},
	}

	start := 0
	for _, item := range items {
		max := rand.Intn(49) + 1
		for i := 0; i < max; i++ {
			item.Serial = fmt.Sprintf("W%06d", i+start)
			serialized, _ := json.Marshal(item)
			if err := stub.PutState(item.Serial, serialized); err != nil {
				return errors.Wrap(err, "failed to put item")
			}
		}
		start += 1000
	}
	return nil
}

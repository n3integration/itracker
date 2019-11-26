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

	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func toValue(v *queryresult.KeyModification) string {
	if v.IsDelete {
		return "null"
	}
	return fmt.Sprintf("%q", string(v.Value))
}

func newError(code int32, format string, args ...interface{}) pb.Response {
	return pb.Response{
		Status:  code,
		Message: fmt.Sprintf(format, args...),
	}
}

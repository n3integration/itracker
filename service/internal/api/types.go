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
package api

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

var opValidationRule = validation.In(fnTransfer, "status").Error("invalid operation")

type (
	Arg    []byte
	Args   [][]byte
	Status int
)

type Item struct {
	Serial       string `json:"serial"`
	Manufacturer string `json:"manufacturer"`
	Code         string `json:"code"`
	Category     string `json:"category"`
	Facility     string `json:"facility"`
	Status       Status `json:"status"`
	SubmittedBy  string `json:"submittedBy"`
}

func (i Item) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Serial, validation.Required),
		validation.Field(&i.Code, validation.Required),
		validation.Field(&i.Manufacturer, validation.Required),
		validation.Field(&i.Facility, validation.Required),
		validation.Field(&i.Category, validation.Required, is.Alphanumeric),
	)
}

type UpdateReq struct {
	Operation string `json:"op"`
	Value     string `json:"value,omitempty"`
}

func (r UpdateReq) Validate() error {
	if r.Operation == fnTransfer {
		return validation.ValidateStruct(&r,
			validation.Field(&r.Operation, validation.Required, opValidationRule),
			validation.Field(&r.Value, validation.Required),
		)
	}
	return validation.ValidateStruct(&r,
		validation.Field(&r.Operation, validation.Required, opValidationRule),
	)
}

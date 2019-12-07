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
	"encoding/json"
	"io"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	apiErrors "github.com/n3integration/itracker/internal/errors"
	"github.com/pkg/errors"
)

type RequestHandler func(req *http.Request) (channel.Response, error)

func Handler(fn RequestHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		v, err := fn(req)
		if err == nil {
			replyWithJson(w, v.Payload)
			return
		}

		switch t := err.(type) {
		case apiErrors.Error:
			w.WriteHeader(t.Code)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		replyWithJson(w, err)
	}
}

func decode(r io.Reader, v validation.Validatable) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return apiErrors.New(http.StatusBadRequest, errors.WithMessage(err, "invalid json request"))
	}

	if err := v.Validate(); err != nil {
		return apiErrors.New(http.StatusBadRequest, err)
	}
	return nil
}

func replyWithJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	switch v := value.(type) {
	case string:
		w.Write([]byte(v))
	case []byte:
		w.Write(v)
	default:
		if content, err := json.Marshal(value); err != nil {
			errContent, _ := json.Marshal(apiErrors.UnknownError)
			w.WriteHeader(apiErrors.UnknownError.Code)
			w.Write(errContent)
		} else {
			w.Write(content)
		}
	}
}

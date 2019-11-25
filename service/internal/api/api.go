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
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

var unknownError = &Error{Status: "error", Message: "an unknown error occurred"}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func decode(w http.ResponseWriter, r io.Reader, v validation.Validatable) (err error) {
	if err = json.NewDecoder(r).Decode(v); err != nil {
		handleBadRequest(w, errors.WithMessage(err, "invalid json request"))
		return
	}

	if err = v.Validate(); err != nil {
		handleBadRequest(w, err)
		return
	}

	return nil
}

func handleBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	replyWithError(w, err)
}

func replyWithError(w http.ResponseWriter, err error) error{
	apiErr := &Error{Status: "error", Message: err.Error()}
	replyWithJson(w, apiErr)
	return err
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
			errContent, _ := json.Marshal(unknownError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errContent)
		} else {
			w.Write(content)
		}
	}
}

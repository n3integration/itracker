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
	"bytes"
	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/n3integration/itracker/internal/app"
	"github.com/n3integration/itracker/internal/logger"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const (
	fnAdd          = "add"
	fnGet          = "get"
	fnGetHistory   = "history"
	fnQuery        = "query"
	fnTransfer     = "transfer"
	fnUpdateStatus = "update"
)

var empty Args

type channelFunc func(request channel.Request, options ...channel.RequestOption) (channel.Response, error)

type InventoryController struct {
	cfg    *app.Config
	client *channel.Client
	sdk    *fabsdk.FabricSDK
}

func NewInventoryController(cfg *app.Config) (*InventoryController, error) {
	logger.Debug("initializing SDK...")
	sdk, err := fabsdk.New(config.FromFile(cfg.ConfigFile))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create SDK")
	}

	logger.Debug("initializing channel context")
	clientContext := sdk.ChannelContext(cfg.ChannelID, fabsdk.WithUser(cfg.UserName), fabsdk.WithOrg(cfg.OrgName))
	client, err := channel.New(clientContext)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to create new channel client")
	}

	logger.Info("fabric client initialized")
	return &InventoryController{
		cfg:    cfg,
		client: client,
		sdk:    sdk,
	}, nil
}

func (c InventoryController) Query(w http.ResponseWriter, _ *http.Request) {
	handle(w, c.newRequest(fnQuery, empty), c.client.Query)
}

func (c InventoryController) Get(w http.ResponseWriter, req *http.Request) {
	handle(w, c.newRequest(fnGet, Args{c.getSerialNumber(req)}), c.client.Query)
}

func (c InventoryController) History(w http.ResponseWriter, req *http.Request) {
	handle(w, c.newRequest(fnGetHistory, Args{c.getSerialNumber(req)}), c.client.Query)
}

func (c InventoryController) Add(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	raw, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleBadRequest(w, errors.WithMessage(err, "bad request"))
		return
	}

	var item Item
	if err := decode(w, bytes.NewReader(raw), &item); err != nil {
		return
	}

	handle(w, c.newRequest(fnAdd, Args{raw}), c.client.Execute)
}

func (c InventoryController) Update(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	var updateReq UpdateReq
	if err := decode(w, req.Body, &updateReq); err != nil {
		return
	}

	switch updateReq.Operation {
	case fnTransfer:
		handle(w, c.newRequest(fnTransfer, Args{c.getSerialNumber(req), Arg(updateReq.Value)}), c.client.Execute)
	default:
		handle(w, c.newRequest(fnUpdateStatus, Args{c.getSerialNumber(req)}), c.client.Execute)
	}
}

func (c InventoryController) Close() error {
	c.sdk.Close()
	return nil
}

func (c InventoryController) getSerialNumber(req *http.Request) Arg {
	return Arg(mux.Vars(req)["serial"])
}

func (c InventoryController) newRequest(fn string, args Args) *channel.Request {
	return &channel.Request{
		ChaincodeID: c.cfg.ChainCodeID,
		Fcn:         fn,
		Args:        args,
	}
}

func handle(w http.ResponseWriter, req *channel.Request, fn channelFunc) {
	response, err := fn(*req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		replyWithError(w, err)
		return
	}

	switch response.ChaincodeStatus {
	case http.StatusOK:
		replyWithJson(w, response.Payload)
	default:
		if response.ChaincodeStatus > 0 {
			w.WriteHeader(int(response.ChaincodeStatus))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		replyWithJson(w, response.Payload)
	}
}

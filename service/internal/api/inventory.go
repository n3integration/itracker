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
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/n3integration/itracker/internal/app"
	apiErrors "github.com/n3integration/itracker/internal/errors"
	"github.com/n3integration/itracker/internal/logger"
	"github.com/pkg/errors"
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

func (c InventoryController) Query(_ *http.Request) (channel.Response, error) {
	return c.client.Query(c.newRequest(fnQuery, empty))
}

func (c InventoryController) Get(req *http.Request) (channel.Response, error) {
	return c.client.Query(c.newRequest(fnGet, Args{c.getSerialNumber(req)}))
}

func (c InventoryController) History(req *http.Request) (channel.Response, error) {
	return c.client.Query(c.newRequest(fnGetHistory, Args{c.getSerialNumber(req)}))
}

func (c InventoryController) Add(req *http.Request) (channel.Response, error) {
	defer req.Body.Close()

	raw, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return channel.Response{}, apiErrors.New(http.StatusBadRequest, errors.WithMessage(err, "bad request"))
	}

	var item Item
	if err := decode(bytes.NewReader(raw), &item); err != nil {
		return channel.Response{}, err
	}

	return c.client.Execute(c.newRequest(fnAdd, Args{raw}))
}

func (c InventoryController) Update(req *http.Request) (channel.Response, error) {
	defer req.Body.Close()

	var updateReq UpdateReq
	if err := decode(req.Body, &updateReq); err != nil {
		return channel.Response{}, err
	}

	switch updateReq.Operation {
	case fnTransfer:
		return c.client.Execute(c.newRequest(fnTransfer, Args{c.getSerialNumber(req), Arg(updateReq.Value)}))
	default:
		return c.client.Execute(c.newRequest(fnUpdateStatus, Args{c.getSerialNumber(req)}))
	}
}

func (c InventoryController) Close() error {
	c.sdk.Close()
	return nil
}

func (c InventoryController) getSerialNumber(req *http.Request) Arg {
	return Arg(mux.Vars(req)["serial"])
}

func (c InventoryController) newRequest(fn string, args Args) channel.Request {
	return channel.Request{
		ChaincodeID: c.cfg.ChainCodeID,
		Fcn:         fn,
		Args:        args,
	}
}

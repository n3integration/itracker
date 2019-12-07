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
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/n3integration/itracker/internal/api"
	"github.com/n3integration/itracker/internal/app"
	"github.com/n3integration/itracker/internal/logger"
	"github.com/n3integration/itracker/internal/middleware"
)

const (
	defaultDir  = "../ui/dist/itracker"
	defaultPort = 8020
)

var (
	chainCodeID string
	channelID   string
	configFile  string
	dir         string
	orgName     string
	port        int
	userName    string
)

func init() {
	flag.StringVar(&chainCodeID, "n", app.DefaultConfig.ChainCodeID, "the chain code id")
	flag.StringVar(&channelID, "C", app.DefaultConfig.ChannelID, "the channel id")
	flag.StringVar(&configFile, "f", app.DefaultConfig.ConfigFile, "the client configuration file")
	flag.StringVar(&orgName, "o", app.DefaultConfig.OrgName, "the organization (msp) name")
	flag.StringVar(&userName, "u", app.DefaultConfig.UserName, "the username")
	flag.StringVar(&dir, "d", defaultDir, "the base directory of the web assets")
	flag.IntVar(&port, "p", defaultPort, "the web server port")
	flag.Parse()
}

func main() {
	inventoryController, err := api.NewInventoryController(app.DefaultConfig)
	if err != nil {
		logger.Fatal("failed to initialize inventory controller: %s", err)
	}

	r := mux.NewRouter()
	r.Use(middleware.Logger)
	apiRouter := r.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/health", api.GetHealth).Methods(http.MethodGet)
	apiRouter.HandleFunc("/inventory", api.Handler(inventoryController.Query)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/inventory", api.Handler(inventoryController.Add)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/inventory/{serial}", api.Handler(inventoryController.Get)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/inventory/{serial}", api.Handler(inventoryController.Update)).Methods(http.MethodPut)
	apiRouter.HandleFunc("/inventory/{serial}/history", api.Handler(inventoryController.History)).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: r,
	}

	logger.Fatal("%s", server.ListenAndServe())
}

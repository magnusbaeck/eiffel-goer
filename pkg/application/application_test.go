// Copyright 2021 Axis Communications AB.
//
// For a full list of individual contributors, please see the commit history.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package application

import (
	"errors"
	"testing"

	"github.com/eiffel-community/eiffel-goer/internal/database"
	"github.com/eiffel-community/eiffel-goer/test/mock_config"
	"github.com/eiffel-community/eiffel-goer/test/mock_database"
	"github.com/eiffel-community/eiffel-goer/test/mock_server"
	"github.com/golang/mock/gomock"
)

// Test that it is possible to get an application.
func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	if app.Config != mockCfg {
		t.Errorf("config not set properly by Get")
	}
	if app.Database == nil {
		t.Error("application did not set up database")
	}
	if app.Router == nil {
		t.Error("application did not set up router")
	}
	if app.Server == nil {
		t.Error("application did not set up server")
	}
}

// Test that it is possible to get an application without a database.
func TestGetNoDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockCfg.EXPECT().GetDBConnectionString().Return("")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	if app.Config != mockCfg {
		t.Errorf("config not set properly by Get")
	}
	if app.Database != nil {
		t.Error("application did not set up database")
	}
	if app.Router == nil {
		t.Error("application did not set up router")
	}
	if app.Server == nil {
		t.Error("application did not set up server")
	}
}

// Test that Get return error if there was an error when getting database.
func TestGetDBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockCfg.EXPECT().GetDBConnectionString().Return("invalid").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")

	_, err := Get(mockCfg)
	if err == nil {
		t.Error("application should have raised error due to invalid database connection string")
	}
}

// Test that getDB return a database interface.
func TestGetDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")

	application := &Application{
		Config: mockCfg,
	}
	db, err := application.getDB()
	if err != nil {
		t.Error(err)
	}
	_, ok := db.(database.Database)
	if !ok {
		t.Error("database from 'getDB' is not a Database interface")
	}
}

// Test that the application creates the v1alpha1 subrouter.
func TestLoadV1Alpha1Routes(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")
	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}

	app.LoadV1Alpha1Routes()
	route := app.Router.Get("v1alpha1")
	if route == nil {
		t.Error("the v1alpha1 route did not get loaded")
	}
}

// Test that the application starts the WebServer & connects to the Database.
func TestStart(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)
	mockServer := mock_server.NewMockServer(ctrl)

	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}

	mockDB.EXPECT().Connect(gomock.Any()).Return(nil)
	mockDB.EXPECT().Close().Return(nil)
	mockServer.EXPECT().WithAddr(":8080").Return(mockServer)
	mockServer.EXPECT().WithRouter(app.Router).Return(mockServer)
	mockServer.EXPECT().Start().Return(nil)
	mockServer.EXPECT().WaitStopped().Return(true)
	mockServer.EXPECT().Error().Return(nil)

	app.Database = mockDB
	app.Server = mockServer

	err = app.Start()
	if err != nil {
		t.Error(err)
	}
}

// Test that the application Start aborts with error if database connect fails
func TestStartAbort(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)
	mockServer := mock_server.NewMockServer(ctrl)

	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}

	mockServer.EXPECT().WithAddr(":8080").Return(mockServer)
	mockServer.EXPECT().WithRouter(app.Router).Return(mockServer)
	mockServer.EXPECT().WaitStopped().Return(true)
	mockServer.EXPECT().Error().Return(nil)
	mockDB.EXPECT().Connect(gomock.Any()).Return(errors.New("did not work"))

	app.Database = mockDB
	app.Server = mockServer

	err = app.Start()
	if err == nil {
		t.Error("application did not abort start after error on database.Connect")
	}
}

// Test that application returns error if server start fails.
func TestStartFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)
	mockServer := mock_server.NewMockServer(ctrl)

	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")
	mockCfg.EXPECT().GetAPIPort().Return("")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}

	mockDB.EXPECT().Connect(gomock.Any()).Return(nil)
	mockDB.EXPECT().Close().Return(nil)
	mockServer.EXPECT().WithAddr("").Return(mockServer)
	mockServer.EXPECT().WithRouter(app.Router).Return(mockServer)
	mockServer.EXPECT().Start().Return(errors.New("error starting"))

	app.Database = mockDB
	app.Server = mockServer

	err = app.Start()
	if err == nil {
		t.Error("application start did not abort when server.Start failed")
	}
}

// Test that application closes the database on Stop.
func TestStop(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	mockCfg.EXPECT().GetDBConnectionString().Return("mongodb://testdb").Times(2)
	mockCfg.EXPECT().GetDatabaseName().Return("testdb")

	app, err := Get(mockCfg)
	if err != nil {
		t.Error(err)
	}

	mockDB.EXPECT().Close().Return(nil)
	app.Database = mockDB
	app.Stop()
}

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
package config

import (
	"os"
	"testing"
)

// Test that it is possible to get a Cfg from Get with values taken from environment variables
func TestGet(t *testing.T) {
	port := "8080"
	connectionString := "connection string"
	databaseName := "database name"
	os.Setenv("CONNECTION_STRING", connectionString)
	os.Setenv("DATABASE_NAME", databaseName)
	os.Setenv("API_PORT", port)

	cfg, ok := Get().(*Cfg)
	if !ok {
		t.Error("cfg returned from get is not a config interface")
	}
	if cfg.connectionString != connectionString {
		t.Error("connection string not set to environment variable CONNECTION_STRING")
	}
	if cfg.databaseName != databaseName {
		t.Error("database name not set to environment variable DATABASE_NAME")
	}
	if cfg.apiPort != port {
		t.Error("api port not set to environment variable API_PORT")
	}
}

// Test that GetDBConnectionString return the connectionString value from the Cfg struct
func TestGetDBConnectionString(t *testing.T) {
	cfg := &Cfg{
		connectionString: "connectionString",
	}
	if cfg.GetDBConnectionString() != "connectionString" {
		t.Error("function does not return the connectionString from Cfg struct")
	}
}

// Test that GetDatabaseName return the databaseName value from the Cfg struct
func TestGetDatabaseName(t *testing.T) {
	cfg := &Cfg{
		databaseName: "databaseName",
	}
	if cfg.GetDatabaseName() != "databaseName" {
		t.Error("function does not return the databaseName from Cfg struct")
	}
}

// Test that GetAPIPort return the value from Cfg struct with a ':' at the start
func TestGetAPIPort(t *testing.T) {
	cfg := &Cfg{
		apiPort: "8080",
	}
	if cfg.GetAPIPort() != ":8080" {
		t.Error("function does not return the apiPort from Cfg struct")
	}
}

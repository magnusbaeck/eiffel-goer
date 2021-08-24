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
	"flag"
	"os"
)

type Config interface {
	GetDBConnectionString() string
	GetDatabaseName() string
	GetAPIPort() string
}

type Cfg struct {
	connectionString string
	databaseName     string

	apiPort string
}

// Parse input parameters to program and return a config with them set.
func Get() Config {
	conf := &Cfg{}

	flag.StringVar(&conf.connectionString, "connectionstring", os.Getenv("CONNECTION_STRING"), "Database connection string.")
	flag.StringVar(&conf.databaseName, "databasename", os.Getenv("DATABASE_NAME"), "Database name.")

	flag.StringVar(&conf.apiPort, "apiport", os.Getenv("API_PORT"), "API port.")

	flag.Parse()
	return conf
}

// Get the connection string for a database.
func (c *Cfg) GetDBConnectionString() string {
	return c.connectionString
}

// Get the name of the database to connect to.
func (c *Cfg) GetDatabaseName() string {
	return c.databaseName
}

// Get API port with a ":" prepended.
func (c *Cfg) GetAPIPort() string {
	return ":" + c.apiPort
}

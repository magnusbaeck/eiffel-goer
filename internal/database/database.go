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
package database

import (
	"context"
	"fmt"
	"net/url"

	"github.com/eiffel-community/eiffel-goer/internal/database/drivers/mongodb"
	"github.com/eiffel-community/eiffel-goer/pkg/schema"
)

type Database interface {
	Connect(context.Context) error
	GetEvents() ([]schema.EiffelEvent, error)
	SearchEvent(string) (schema.EiffelEvent, error)
	UpstreamDownstreamSearch(string) ([]schema.EiffelEvent, error)
	GetEventByID(string) (schema.EiffelEvent, error)
	Close() error
}

// Get a new Database.
func Get(connectionString string, databaseName string) (Database, error) {
	db, err := get(connectionString, databaseName)
	if err != nil {
		return nil, err
	}
	var database Database = db
	return database, nil
}

// Get a database driver based on the connectionURL scheme supplied in the configuration.
func get(connectionString string, databaseName string) (Database, error) {
	connectionURL, err := url.Parse(connectionString)
	if err != nil {
		return nil, err
	}
	switch connectionURL.Scheme {
	case "mongodb":
		return mongodb.Get(connectionString, databaseName)
	}
	return nil, fmt.Errorf("cannot find database for scheme '%s'", connectionURL.Scheme)
}

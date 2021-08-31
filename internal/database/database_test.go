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
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGet(t *testing.T) {
	_, err := Get("mongodb://db/test", &log.Entry{})
	if err != nil {
		t.Error(err)
	}
}

func TestGetUnknownScheme(t *testing.T) {
	_, err := Get("unknown://db/test", &log.Entry{})
	if err == nil {
		t.Error("possible to get a database with unknown:// scheme")
	}
}

func TestGetUnparsableScheme(t *testing.T) {
	_, err := Get("://", &log.Entry{})
	if err == nil {
		t.Error("possible to get a database with an unparsable scheme")
	}
}

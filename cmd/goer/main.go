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
package main

import (
	"context"

	"github.com/eiffel-community/eiffel-goer/internal/config"
	"github.com/eiffel-community/eiffel-goer/internal/logger"
	"github.com/eiffel-community/eiffel-goer/pkg/application"
)

// Start up the Goer application.
func main() {
	cfg := config.Get()
	ctx := context.Background()

	app, err := application.Get(cfg)
	if err != nil {
		logger.Error.Panic(err)
	}

	app.LoadV1Alpha1Routes()

	logger.Debug.Println("Starting up.")
	err = app.Start(ctx)
	if err != nil {
		logger.Error.Panic(err)
	}
}

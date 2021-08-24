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
package logger

import (
	"log"
	"os"
)

// Add some logging handlers that can be used to append
// log level to the log message. Also redirect Error to Stderr.
var (
	Debug   = log.New(os.Stdout, "DEBUG\t", log.Ldate|log.Ltime)
	Info    = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	Warning = log.New(os.Stdout, "WARNING\t", log.Ldate|log.Ltime)
	Error   = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)

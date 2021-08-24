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
package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server interface {
	WithAddr(string) Server
	WithErrLogger(*log.Logger) Server
	WithRouter(*mux.Router) Server
	Start() error
	Error() error
	WaitRunning() bool
	WaitStopped() bool
	Close() error
}

type WebServer struct {
	server  *http.Server
	running chan bool
	stopped chan bool
	err     error
}

// Create a new WebServer.
func Get() Server {
	return &WebServer{
		server:  &http.Server{},
		running: make(chan bool, 2), // Buffer up to two messages.
		stopped: make(chan bool, 2), // Buffer up to two messages.
	}
}

// Add an address to the server.
func (s *WebServer) WithAddr(addr string) Server {
	s.server.Addr = addr
	return s
}

// Add an error logger to the server.
func (s *WebServer) WithErrLogger(l *log.Logger) Server {
	s.server.ErrorLog = l
	return s
}

// Add a router to the server.
func (s *WebServer) WithRouter(router *mux.Router) Server {
	s.server.Handler = router
	return s
}

// Start the webserver.
func (s *WebServer) Start() error {
	if len(s.server.Addr) == 0 {
		s.running <- false
		return errors.New("server missing address")
	}
	if s.server.Handler == nil {
		s.running <- false
		return errors.New("server missing handler")
	}
	go s.run()
	return nil
}

// Get error message from the webserver.
func (s *WebServer) Error() error {
	return s.err
}

// Wait for the webserver to start running.
// Note that this will send 'true' first, and then 'false' if the server crashes.
func (s *WebServer) WaitRunning() bool {
	return <-s.running
}

// Wait for the webserver to stop.
func (s *WebServer) WaitStopped() bool {
	return <-s.stopped
}

// Run the webserver.
func (s *WebServer) run() {
	s.running <- true
	err := s.server.ListenAndServe()
	s.err = err
	s.stopped <- true
	s.running <- false
}

// Close the webserver.
func (s *WebServer) Close() error {
	err := s.server.Close()
	return err
}

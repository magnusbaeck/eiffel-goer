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

// This package tests that all endpoints are added and respond to requests.
// The function of each handler is tested separately in each handler.
package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eiffel-community/eiffel-goer/pkg/application"
	"github.com/eiffel-community/eiffel-goer/pkg/schema"
	"github.com/eiffel-community/eiffel-goer/test/mock_config"
	"github.com/eiffel-community/eiffel-goer/test/mock_database"
	"github.com/golang/mock/gomock"
)

// Test that the endpoint /events for a single event is added properly.
func TestAddedRoutesEventRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	eventID := "3fabaa6b-5343-4d74-8af9-dc2e4c1f2827"

	mockCfg.EXPECT().GetDBConnectionString().Return("")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")
	mockDB.EXPECT().GetEventByID(eventID).Return(schema.EiffelEvent{}, nil)

	app, err := application.Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	app.Database = mockDB
	app.LoadV1Alpha1Routes()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1alpha1/events/"+eventID, nil)
	app.Router.ServeHTTP(responseRecorder, request)
	expectedStatusCode := http.StatusOK
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/v1alpha1/events/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
}

// Test that the endpoint /events for a many events is added properly.
func TestAddedRoutesEventReadAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	mockCfg.EXPECT().GetDBConnectionString().Return("")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")
	mockDB.EXPECT().GetEvents().Return([]schema.EiffelEvent{}, nil)

	app, err := application.Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	app.Database = mockDB
	app.LoadV1Alpha1Routes()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1alpha1/events", nil)
	app.Router.ServeHTTP(responseRecorder, request)
	expectedStatusCode := http.StatusNotImplemented
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/v1alpha1/events', got '%d'", expectedStatusCode, responseRecorder.Code)
	}
}

// Test that the endpoint /search is added properly.
func TestAddedRoutesSearchRead(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	eventID := "3fabaa6b-5343-4d74-8af9-dc2e4c1f2827"

	// Setting DBConnection string to "" so that application.Get does not
	// attempt to create its own database connection. After initializing
	// the application, add the MockDB
	mockCfg.EXPECT().GetDBConnectionString().Return("")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")
	mockDB.EXPECT().SearchEvent("id").Return(schema.EiffelEvent{}, nil)

	app, err := application.Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	app.Database = mockDB
	app.LoadV1Alpha1Routes()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/v1alpha1/search/"+eventID, nil)
	app.Router.ServeHTTP(responseRecorder, request)
	expectedStatusCode := http.StatusNotImplemented
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/v1alpha1/search/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
}

// Test that the post endpoint /search is added properly.
func TestAddedRoutesSearchUpstreamDownstream(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	eventID := "3fabaa6b-5343-4d74-8af9-dc2e4c1f2827"

	// Setting DBConnection string to "" so that application.Get does not
	// attempt to create its own database connection. After initializing
	// the application, add the MockDB
	mockCfg.EXPECT().GetDBConnectionString().Return("")
	mockCfg.EXPECT().GetAPIPort().Return(":8080")
	mockDB.EXPECT().UpstreamDownstreamSearch("id").Return([]schema.EiffelEvent{}, nil)

	app, err := application.Get(mockCfg)
	if err != nil {
		t.Error(err)
	}
	app.Database = mockDB
	app.LoadV1Alpha1Routes()

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/v1alpha1/search/"+eventID, nil)
	app.Router.ServeHTTP(responseRecorder, request)
	expectedStatusCode := http.StatusNotImplemented
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/v1alpha1/search/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
}

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
package events

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eiffel-community/eiffel-goer/pkg/schema"
	"github.com/eiffel-community/eiffel-goer/test/mock_config"
	"github.com/eiffel-community/eiffel-goer/test/mock_database"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

var activityJSON = []byte(`
{
    "data": {
        "name": "Test activity"
    },
    "links": [],
    "meta": {
        "id": "e04cf9d3-4d57-471e-bd65-f8fc20d21d84",
        "time": 1629449650361,
        "type": "EiffelActivityTriggeredEvent",
        "version": "3.0.0"
    }
}
`)

func loadEvent() (schema.EiffelEvent, error) {
	event := schema.EiffelEvent{}
	err := json.Unmarshal(activityJSON, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func TestRead(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	event, err := loadEvent()
	if err != nil {
		t.Error(err)
	}
	eventID := event.Meta.ID
	mockDB.EXPECT().GetEventByID(eventID).Return(event, nil)

	app := Get(mockCfg, mockDB)
	handler := mux.NewRouter()
	handler.HandleFunc("/events/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", app.Read)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/events/"+eventID, nil)

	handler.ServeHTTP(responseRecorder, request)

	expectedStatusCode := http.StatusOK
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/events/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
	eventFromResponse := schema.EiffelEvent{}
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &eventFromResponse)
	if err != nil {
		t.Error(err)
	}
	if eventFromResponse.Meta.ID != event.Meta.ID {
		t.Error("event returned with response is not the same as the one in DB")
	}
}

func TestReadBadQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	event, err := loadEvent()
	if err != nil {
		t.Error(err)
	}
	eventID := event.Meta.ID
	mockDB.EXPECT().GetEventByID(eventID).Return(event, nil)

	app := Get(mockCfg, mockDB)
	handler := mux.NewRouter()
	handler.HandleFunc("/events/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", app.Read)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/events/"+eventID, nil)
	q := request.URL.Query()
	q.Add("nah", "hello")
	request.URL.RawQuery = q.Encode()

	handler.ServeHTTP(responseRecorder, request)

	expectedStatusCode := http.StatusBadRequest
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/events/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
}

func TestReadEventNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockCfg := mock_config.NewMockConfig(ctrl)
	mockDB := mock_database.NewMockDatabase(ctrl)

	event, err := loadEvent()
	if err != nil {
		t.Error(err)
	}
	eventID := event.Meta.ID
	mockDB.EXPECT().GetEventByID(eventID).Return(schema.EiffelEvent{}, errors.New("does not exist"))

	app := Get(mockCfg, mockDB)
	handler := mux.NewRouter()
	handler.HandleFunc("/events/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}", app.Read)

	responseRecorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/events/"+eventID, nil)

	handler.ServeHTTP(responseRecorder, request)

	expectedStatusCode := http.StatusNotFound
	if responseRecorder.Code != expectedStatusCode {
		t.Errorf("Want status '%d' for '/events/%s', got '%d'", expectedStatusCode, eventID, responseRecorder.Code)
	}
}

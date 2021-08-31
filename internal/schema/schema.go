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

// This package shall be removed when the Eiffel Go SDK is finished.
package schema

type EiffelEvent struct {
	Meta  Meta        `json:"meta"`
	Links []Link      `json:"links"`
	Data  interface{} `json:"data"`
}

type Link struct {
	Target string `json:"target"`
	Type   string `json:"type"`
}

type Meta struct {
	ID       string   `json:"id"`
	Security Security `json:"security,omitempty"`
	Source   Source   `json:"source,omitempty"`
	Tags     []Tag    `json:"tags,omitempty"`
	Time     int      `json:"time"`
	Type     string   `json:"type"`
	Version  string   `json:"version"`
}

type Security struct {
	AuthorIdentity      string                   `json:"authorIdentity,omitempty"`
	IntegrityProtection IntegrityProtection      `json:"integrityProtection"`
	SequenceProtection  []SequenceProtectionItem `json:"sequenceProtection,omitempty"`
}

type Source struct {
	DomainID   string `json:"domainId,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Serializer string `json:"serializer,omitempty"`
	URI        string `json:"uri,omitempty"`
}

type Tag string

type IntegrityProtection struct {
	Alg       string `json:"alg"`
	PublicKey string `json:"publicKey,omitempty"`
	Signature string `json:"signature"`
}

type SequenceProtectionItem struct {
	Position     int    `json:"position"`
	SequenceName string `json:"sequenceName"`
}

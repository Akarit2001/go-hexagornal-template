// internal/transport/jsonapi/generic.go
package jsonapi

import (
	"encoding/json"
)

// -----------------------------------------------------------------------------
// Reusable infrastructure
// -----------------------------------------------------------------------------

type Links struct {
	Self    string `json:"self,omitempty"`
	Related string `json:"related,omitempty"`
}

type ResourceIdentifier struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Relationship can hold a single linkage, an array, or be empty.
type Relationship struct {
	Links *Links          `json:"links,omitempty"`
	Data  json.RawMessage `json:"data,omitempty"` // either ResourceIdentifier, []ResourceIdentifier, or null
	Meta  json.RawMessage `json:"meta,omitempty"`
}

// -----------------------------------------------------------------------------
// Resource and Document structures
// -----------------------------------------------------------------------------

// Resource is a JSON:API resource object whose Attributes are user-defined.
type Resource[Attr any] struct {
	Type          string                  `json:"type"`
	ID            string                  `json:"id,omitempty"`
	Attributes    Attr                    `json:"attributes"`
	Relationships map[string]Relationship `json:"relationships,omitempty"`
	Links         *Links                  `json:"links,omitempty"`
	Meta          json.RawMessage         `json:"meta,omitempty"`
}

// ErrorObject represents a JSON:API error object.
type ErrorObject struct {
	ID     string          `json:"id,omitempty"`
	Status string          `json:"status,omitempty"`
	Code   string          `json:"code,omitempty"`
	Title  string          `json:"title,omitempty"`
	Detail string          `json:"detail,omitempty"`
	Source json.RawMessage `json:"source,omitempty"`
	Links  *Links          `json:"links,omitempty"`
	Meta   json.RawMessage `json:"meta,omitempty"`
}

// JSONAPI represents the optional top-level jsonapi object.
type JSONAPI struct {
	Version string          `json:"version,omitempty"` // usually "1.1"
	Meta    json.RawMessage `json:"meta,omitempty"`
}

// Document represents a generic JSON:API document that can be a single resource,
// a collection, null data, errors, and includes compound documents.
type Document[Attr any] struct {
	Data     json.RawMessage  `json:"data,omitempty"`     // object, array, or null
	Errors   []ErrorObject    `json:"errors,omitempty"`   // mutually exclusive with Data
	Included []Resource[Attr] `json:"included,omitempty"` // compound documents
	Links    *Links           `json:"links,omitempty"`
	Meta     json.RawMessage  `json:"meta,omitempty"`
	JSONAPI  *JSONAPI         `json:"jsonapi,omitempty"`
}

// -----------------------------------------------------------------------------
// Optional helpers for common cases
// -----------------------------------------------------------------------------

// NewSingle creates a Document with a single resource as data.
func NewSingle[Attr any](r Resource[Attr]) Document[Attr] {
	raw, _ := json.Marshal(r)
	return Document[Attr]{Data: raw}
}

// NewCollection creates a Document with a collection of resources as data.
func NewCollection[Attr any](rs []Resource[Attr]) Document[Attr] {
	raw, _ := json.Marshal(rs)
	return Document[Attr]{Data: raw}
}

// NewNullData creates a Document with `data: null`.
func NewNullData[Attr any]() Document[Attr] {
	return Document[Attr]{Data: []byte("null")}
}

// NewErrors creates a Document with errors.
func NewErrors(msgs ...ErrorObject) Document[any] {
	return Document[any]{Errors: msgs}
}

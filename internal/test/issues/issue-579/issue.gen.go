// Package issue579 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/reedom/oapi-codegen version (devel) DO NOT EDIT.
package issue579

import (
	openapi_types "github.com/reedom/oapi-codegen/pkg/types"
)

// AliasedDate defines model for AliasedDate.
type AliasedDate = openapi_types.Date

// Pet defines model for Pet.
type Pet struct {
	Born   *AliasedDate        `json:"born,omitempty"`
	BornAt *openapi_types.Date `json:"born_at,omitempty"`
}

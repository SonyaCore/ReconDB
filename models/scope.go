package models

// Scopes struct model for scope api
type Scopes struct {
	// CompanyName is a identifier for detecting if company registers is company endpoint.
	CompanyName string `json:"company_name"`

	// ScopeType is the type of scope.
	ScopeType string `json:"scope_type"`

	// Scope is the scope of the program with specified ScopeType
	Scope string `json:"scope"`
}

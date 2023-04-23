package models

// Scopes struct model for scope api
// CompanyName is a identifier for detecting if company registers is company endpoint.
// ScopeType is the type of scope which are 4 of them : [cidr , single , ip , wildcard]
// Scope is the scope of the program with specified ScopeType
type Scopes struct {
	CompanyName string `json:"company_name"`
	ScopeType   string `json:"scope_type"`
	Scope       string `json:"scope"`
}

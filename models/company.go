package models

// Company struct model for company api
type Company struct {
	// CompanyName is the name of company for setting scopes within that company
	CompanyName string `json:"company_name"`

	// ProgramType is the type of current company program
	// there are only two types for that [vdp,rdp].
	ProgramType string `json:"program_type"`
}

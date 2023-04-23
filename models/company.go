package models

// Company struct model for company api
// CompanyName is the name of company for setting scopes within that company
// ProgramType is the type of current company program there are only two types for that [vdp,rdp].
type Company struct {
	CompanyName string `json:"company_name"`
	ProgramType string `json:"program_type"`
}

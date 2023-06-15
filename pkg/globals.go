package pkg

// Scopes validate types for inserting in scope
// in asset only valid types are [ip , single]
var Scopes = []string{
	"wildcard", "cidr", "single", "ip",
}

// ProgramTypes
// []string used to check company program type
var ProgramTypes = []string{
	"vdp", "rdp",
}

// WildCardPattern
// a regex to validate wildcard domain
var WildCardPattern = `^(\*|(\*|\*\.)?(\*|\*\.)\w+(\.\w+)*(\.\*|\*)?(\.\*|\*)?)$`

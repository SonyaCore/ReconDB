package middlewares

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
var WildCardPattern = `^(\*|(\*|\*\.)?\w+(\.\w+)*(\.\*|\*)?)$`

// IPPattern for valdiatin ip with regex
// not used for now
var IPPattern = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`

// DomainPattern
// validating domain names + port with regex
var DomainPattern = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})(:[0-9]{1,5})?$`

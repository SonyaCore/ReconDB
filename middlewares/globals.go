package middlewares

var Scopes = []string{
	"wildcard", "cidr", "single", "ip",
}

var ProgramTypes = []string{
	"vdp", "rdp",
}

var WildCardPattern = `^(\*|(\*|\*\.)?\w+(\.\w+)*(\.\*|\*)?)$`

var DomainPattern = `^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})(:[0-9]{1,5})?$`

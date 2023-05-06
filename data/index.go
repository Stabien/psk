package data

var Stacks = map[string]string{
	"React.js": "reactjs",
	"Node.js":  "nodejs",
}

var StackOptions = map[string]map[string][]string{
	"reactjs": {},
	"nodejs": {
		"language": {"JavaScript", "TypeScript"},
	},
}

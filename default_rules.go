package vocx

const defaultRules = `
{
	"letters": {
		"a": "a",
		"b": "b",
		"c": "ts",
		"ĉ": "cz",
		"d": "d",
		"e": "e",
		"f": "f",
		"g": "g",
		"ĝ": "dż",
		"h": "h",
		"ĥ": "ch",
		"i": "ij",
		"j": "y",
		"ĵ": "rz",
		"k": "k",
		"l": "l",
		"m": "m",
		"n": "n",
		"o": "o",
		"p": "p",
		"r": "r",
		"s": "s",
		"ŝ": "sz",
		"t": "t",
		"u": "u",
		"ŭ": "ł",
		"v": "w",
		"z": "z"
	},
	"fragments": [
		{ "match": "atsij", "replace": "atssij" },
		{ "match": "ide\b", "replace": "ijde" },
		{ "match": "io\b", "replace": "ijo" },
		{ "match": "ioy\b", "replace": "ijoj" },
		{ "match": "ioyn\b", "replace": "ijojn" },
		{ "match": "feyo\b", "replace": "fejo" },
		{ "match": "feyoy\b", "replace": "feyoj" },
		{ "match": "feyoyn\b", "replace": "feyoj" },
		{ "match": "^ekzij", "replace": "ekzji" },
		{ "match": "tssijl", "replace": "tssil" },
		{ "match": "ijuy", "replace": "iuyy" },
		{ "match": "ijeh", "replace": "ije" },
		{ "match": "sijlo", "replace": "ssilo" }
	],
	"overrides": [
		{ "eo": "ok", "pl": "ohk" }
	]
}
`

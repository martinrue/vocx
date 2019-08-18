package vocx_test

import (
	"fmt"
	"testing"

	"github.com/martinrue/vocx"
)

func TestTranscribeWithDefaultRules(t *testing.T) {
	tests := []struct {
		Input    string
		Expected string
	}{
		{"Saluton", "saluton"},
		{"Saluton.", "saluton."},
		{"Saluton, kiel vi fartas?", "saluton, kijel wij fartas?"},
		{"La oka numero estas ok.", "la oka numero estas ohk."},
		{"Tiel estas la mondo.", "tiel estas la mondo."},
		{"La internacia lingvo estas tre facila.", "la ijnternatssija lijngwo estas tre fatssila."},
		{"abcĉdefgĝhĥijĵklmnoprsŝtuŭvz", "abtssczdefgdżhchijyrzklmnoprssztułwz"},
	}

	transcriber := vocx.NewTranscriber()

	for i, test := range tests {
		t.Run(fmt.Sprintf("default-rules-test-%d", i), func(t *testing.T) {
			actual := transcriber.Transcribe(test.Input)

			if actual != test.Expected {
				t.Fatalf("Expected [%s], got [%s]", test.Expected, actual)
			}
		})
	}
}

func TestTranscribeWithCustomRules(t *testing.T) {
	customRules := `
		{
			"letters": {
				"a": "a",
				"b": "b",
				"c": "tss",
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
				{ "match": "ci\b", "replace": "cyjx" },
				{ "match": "ide\b", "replace": "ijdex" },
				{ "match": "io\b", "replace": "ijox" },
				{ "match": "ioy\b", "replace": "ijojx" },
				{ "match": "ioyn\b", "replace": "ijojnx" },
				{ "match": "feyo\b", "replace": "fejox" },
				{ "match": "feyoy\b", "replace": "feyojx" },
				{ "match": "feyoyn\b", "replace": "feyojx" },
				{ "match": "^tij", "replace": "tix" },
				{ "match": "^ekzij", "replace": "ekzjix" }
			],
			"overrides": [
				{ "eo": "ok", "pl": "ohkx" }
			]
		}
	`

	tests := []struct {
		Input    string
		Expected string
	}{
		{"Saluton", "saluton"},
		{"Saluton.", "saluton."},
		{"Saluton, kiel vi fartas?", "saluton, kijel wij fartas?"},
		{"La oka numero estas ok.", "la oka numero estas ohkx."},
		{"Tiel estas la mondo.", "tixel estas la mondo."},
		{"La internacia lingvo estas tre facila.", "la ijnternatssija lijngwo estas tre fatssijla."},
		{"abcĉdefgĝhĥijĵklmnoprsŝtuŭvz", "abtssczdefgdżhchijyrzklmnoprssztułwz"},
	}

	transcriber := vocx.NewTranscriber()
	transcriber.LoadRules(customRules)

	for i, test := range tests {
		t.Run(fmt.Sprintf("custom-rules-test-%d", i), func(t *testing.T) {
			actual := transcriber.Transcribe(test.Input)

			if actual != test.Expected {
				t.Fatalf("Expected [%s], got [%s]", test.Expected, actual)
			}
		})
	}
}

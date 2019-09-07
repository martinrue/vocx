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
		{"Tiel estas la mondo.", "tijel estas la mondo."},
		{"Mi ricevis mesaĝon", "mij rijtsewijs mesadżon"},
		{"La internacia lingvo estas tre facila.", "la ijnternatssija lijngwo estas tre fatssila."},
		{"Abcĉdefgĝhĥijĵklmnoprsŝtuŭvz", "abtsczdefgdżhchijyrzklmnoprssztułwz"},
		{"La sistemo simple estas bonega", "la syystemo syymple estas bonega"},
		{"Saluton s-ro kaj s-ino", "saluton sjijnjoro kay sjijnjorijno"},
		{"Okazas, okazis, okazos", "okazas, okazyjs, okazos"},
		{"cxgxhxjxsxux", "czdżchrzszł"},
		{"ktp", "ko-to-po"},
		{"k.t.p", "ko-to-po"},
		{"atm", "antałtagmeze"},
		{"ptm", "posttagmeze"},
		{"bv", "bonvolu"},
		{"1", "unu"},
		{"5", "kvijn"},
		{"7,1", "sep, komo unu"},
		{"7,15", "sep, komo dek kvijn"},
		{"12", "dek du"},
		{"100", "tsent"},
		{"600", "ses tsent"},
		{"110", "tsent dek"},
		{"115", "tsent dek kvijn"},
		{"259", "du tsent kvijn dek nał"},
		{"999", "nał tsent nał dek nał"},
		{"1150", "mijl, tsent kvijn dek"},
		{"1250", "mijl, du tsent kvijn dek"},
		{"1.268", "mijl, du tsent ses dek ohk"},
		{"5.233,55", "kvijn mijl, du tsent trij dek trij, komo kvijn dek kvijn"},
		{"839241,12", "ohk tsent trij dek nał mijl, du tsent kvar dek unu, komo dek du"},
		{"1000000", "mijlijono"},
		{"2000000", "du mijlijono"},
		{"9.500123", "nał mijlijono, kvijn tsent mijl, tsent du dek trij"},
		{"249.500123", "du tsent kvar dek nał mijlijono, kvijn tsent mijl, tsent du dek trij"},
	}

	transcriber := vocx.NewTranscriber()

	for i, test := range tests {
		t.Run(fmt.Sprintf("default-rules-test-%d", i+1), func(t *testing.T) {
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
		{"Mi ricevis mesaĝon", "mij rijtsewijs mesadżon"},
		{"La internacia lingvo estas tre facila.", "la ijnternatssija lijngwo estas tre fatssijla."},
		{"Abcĉdefgĝhĥijĵklmnoprsŝtuŭvz", "abtsczdefgdżhchijyrzklmnoprssztułwz"},
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

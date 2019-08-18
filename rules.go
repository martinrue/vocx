package vocx

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// Override is used to replace a full word.
type Override struct {
	Esperanto string `json:"eo"`
	Polish    string `json:"pl"`
}

// Fragment is used to replace a fragment of a word.
type Fragment struct {
	Match   string `json:"match"`
	Replace string `json:"replace"`
}

// Rules capture the set of rules for a transcription.
type Rules struct {
	Letters   map[string]string `json:"letters"`
	Fragments []Fragment        `json:"fragments"`
	Overrides []Override        `json:"overrides"`
}

func (r *Rules) findOverride(word string) string {
	punctuation := `,.!?;:*"'«»„“<>()[]{}`

	prefix := string(word[0])
	prefixed := strings.Contains(punctuation, prefix)

	if prefixed {
		word = word[1:]
	}

	postfix := string(word[len(word)-1])
	postfixed := strings.Contains(punctuation, postfix)

	if postfixed {
		word = word[0 : len(word)-1]
	}

	word = strings.ToLower(word)

	for _, override := range r.Overrides {
		if strings.ToLower(override.Esperanto) == word {
			return fmt.Sprintf("%s%s%s", iff(prefixed, prefix), override.Polish, iff(postfixed, postfix))
		}
	}

	return ""
}

func loadRules(data string) (*Rules, error) {
	r := &Rules{}

	if err := json.NewDecoder(strings.NewReader(data)).Decode(&r); err != nil {
		return nil, err
	}

	for _, fragment := range r.Fragments {
		if _, err := regexp.Compile(fragment.Match); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func iff(expr bool, value string) string {
	if expr {
		return value
	}

	return ""
}

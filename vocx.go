package vocx

import (
	"fmt"
	"regexp"
	"strings"
)

// Transcriber handles transcribing text.
type Transcriber struct {
	rules *Rules
}

// Transcribe transcribes text using the current rules.
func (t *Transcriber) Transcribe(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")

	words := transform(strings.Split(text, " "), func(word string) string {
		word = strings.TrimSpace(word)

		if word == "" {
			return word
		}

		if override := t.rules.findOverride(word); override != "" {
			return override
		}

		letters := strings.Split(word, "")

		for i, letter := range letters {
			l, ok := t.rules.Letters[strings.ToLower(letter)]

			if !ok {
				continue
			}

			letters[i] = l
		}

		word = strings.Join(letters, "")

		for _, fragment := range t.rules.Fragments {
			r := regexp.MustCompile(fragment.Match)
			word = r.ReplaceAllString(word, fragment.Replace)
		}

		return word
	})

	return strings.Join(words, " ")
}

// LoadRules loads a new set of rules from the JSON string.
func (t *Transcriber) LoadRules(json string) {
	rules, err := loadRules(json)
	if err != nil {
		panic(fmt.Errorf("failed to load rules: %v", err))
	}

	t.rules = rules
}

// NewTranscriber returns a transcriber with default rules.
func NewTranscriber() *Transcriber {
	t := &Transcriber{}
	t.LoadRules(defaultRules)
	return t
}

func transform(items []string, fn func(item string) string) []string {
	transformed := make([]string, 0)

	for _, item := range items {
		str := fn(item)

		if str != "" {
			transformed = append(transformed, str)
		}
	}

	return transformed
}

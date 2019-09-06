package vocx

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Transcriber handles transcribing text.
type Transcriber struct {
	rules *Rules
}

// Transcribe transcribes text using the current rules.
func (t *Transcriber) Transcribe(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")

	isNumber := func(word string) bool {
		word = strings.ReplaceAll(strings.ReplaceAll(word, ",", ""), ".", "")
		_, err := strconv.ParseFloat(word, 64)
		return err == nil
	}

	parseNumber := func(word string) (int64, int64) {
		parts := strings.Split(strings.ReplaceAll(word, ".", ""), ",")

		if len(parts) == 1 {
			whole, _ := strconv.ParseInt(parts[0], 10, 64)
			return whole, 0
		}

		if len(parts) == 2 {
			whole, _ := strconv.ParseInt(parts[0], 10, 64)
			fraction, _ := strconv.ParseInt(parts[1], 10, 64)
			return whole, fraction
		}

		return 0, 0
	}

	words := transform(strings.Split(text, " "), func(word string) string {
		word = strings.TrimSpace(word)

		if word == "" {
			return word
		}

		if isNumber(word) {
			return t.transcribeNumber(parseNumber(word))
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
func (t *Transcriber) LoadRules(json string) error {
	rules, err := loadRules(json)
	if err != nil {
		return err
	}

	t.rules = rules
	return nil
}

func (t *Transcriber) transcribeNumber(whole int64, fraction int64) string {
	result := t.transcribeNumberPart(whole)

	if fraction != 0 && fraction < 100 {
		return fmt.Sprintf("%s, komo %s", result, t.transcribeNumberPart(fraction))
	}

	return result
}

func (t *Transcriber) transcribeNumberPart(number int64) string {
	getNumber := func(number int64) string {
		return t.rules.Numbers[strconv.FormatInt(number, 10)]
	}

	transcribe := func(one, ten, hundred int64, includeOne bool) string {
		result := ""

		if hundred > 0 {
			if hundred == 1 {
				result += fmt.Sprintf(" %s", getNumber(100))
			} else {
				result += fmt.Sprintf(" %s %s", getNumber(hundred), getNumber(100))
			}
		}

		if ten > 0 {
			if ten == 1 {
				result += fmt.Sprintf(" %s", getNumber(10))
			} else {
				result += fmt.Sprintf(" %s %s", getNumber(ten), getNumber(10))
			}
		}

		if one == 1 {
			if includeOne {
				result += fmt.Sprintf(" %s", getNumber(1))
			}
		} else if one != 0 {
			result += fmt.Sprintf(" %s", getNumber(one))
		}

		return strings.TrimSpace(result)
	}

	ones := number % 10
	tens := (number / 10) % 10
	hundreds := (number / 100) % 10
	thousands := (number / 1000) % 10
	tenThousands := (number / 10000) % 10
	hundredThousands := (number / 100000) % 10
	millions := (number / 1000000) % 10
	tenMillions := (number / 10000000) % 10
	hundredMillions := (number / 100000000) % 10

	first := transcribe(ones, tens, hundreds, true)
	second := transcribe(thousands, tenThousands, hundredThousands, false)
	third := transcribe(millions, tenMillions, hundredMillions, false)

	hasValues := func(n1, n2, n3 int64) bool {
		return n1 != 0 || n2 != 0 || n3 != 0
	}

	result := ""

	if hasValues(millions, tenMillions, hundredMillions) {
		result += fmt.Sprintf("%s %s,", third, getNumber(1000000))
	}

	if hasValues(thousands, tenThousands, hundredThousands) {
		result += fmt.Sprintf(" %s %s,", second, getNumber(1000))
	}

	if hasValues(ones, tens, hundreds) {
		result += fmt.Sprintf(" %s", first)
	}

	return strings.TrimRight(strings.TrimSpace(result), ",")
}

// NewTranscriber returns a transcriber with default rules.
func NewTranscriber() *Transcriber {
	t := &Transcriber{}
	_ = t.LoadRules(defaultRules)
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

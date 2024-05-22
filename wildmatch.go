// Package wildmatch used to match strings against a simple wildcard pattern.
// Tests a wildcard pattern `p` against an input string `s`. Returns true only when `p` matches the entirety of `s`.
//
// See also the example described on [wikipedia](https://en.wikipedia.org/wiki/Matching_wildcards) for matching wildcards.
//
// No escape characters are defined.
//
// - `?` matches exactly one occurrence of any character.
// - `*` matches arbitrary many (including zero) occurrences of any character.
//
// Examples matching wildcards:
// ``` go
// import "github.com/becheran/wildmatch-go"
// wildmatch.NewWildMatch("cat").Matches("cat")
// wildmatch.NewWildMatch("*cat*").Matches("dog_cat_dog")
// wildmatch.NewWildMatch("c?t").Matches("cat")
// wildmatch.NewWildMatch("c?t").Matches("cot")
// ```
// Examples not matching wildcards:
// ``` go
// import "github.com/becheran/wildmatch-go"
// wildmatch.NewWildMatch("dog").Matches("cat")
// wildmatch.NewWildMatch("*d").Matches("cat")
// wildmatch.NewWildMatch("????").Matches("cat")
// wildmatch.NewWildMatch("?").Matches("cat")
// ```
package wildmatch

/// WildMatch is a wildcard matcher used to match strings.
type WildMatch struct {
	pattern string
}

func (w *WildMatch) String() string {
	return w.pattern
}

// NewWildMatch creates new pattern matcher.
func NewWildMatch(pattern string) *WildMatch {
	simplified := ""
	prevWasStar := false
	for _, currentChar := range pattern {
		if currentChar == '*' {
			if !prevWasStar {
				simplified += "*"
			}
			prevWasStar = true
		} else {
			prevWasStar = false
			simplified += string(currentChar)
		}
	}
	return &WildMatch{
		pattern: simplified,
	}
}

// Matches indicates whether the matcher finds a match in the input string.
func (w *WildMatch) Matches(input string) bool {
	if w.pattern == "" {
		return input == ""
	}

	idxInput := 0
	idxPattern := 0
	idxStart := -1
	idxMatch := 0
	for idxInput < len(input) {
		if idxPattern < len(w.pattern) && w.pattern[idxPattern] == '*' {
			idxStart = idxPattern
			idxMatch = idxInput
			idxPattern++
		} else if idxPattern < len(w.pattern) && (w.pattern[idxPattern] == '?' || w.pattern[idxPattern] == input[idxInput]) {
			idxInput++
			idxPattern++
		} else if idxStart != -1 {
			idxPattern = idxStart + 1
			idxMatch++
			idxInput = idxMatch
		} else {
			return false
		}
	}

	for idxPattern < len(w.pattern) && w.pattern[idxPattern] == '*' {
		idxPattern++
	}

	return idxPattern == len(w.pattern)
}

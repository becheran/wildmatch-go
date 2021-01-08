package wildmatch

import "strings"

//! Match strings against a simple wildcard pattern.
//! Tests a wildcard pattern `p` against an input string `s`. Returns true only when `p` matches the entirety of `s`.
//!
//! See also the example described on [wikipedia](https://en.wikipedia.org/wiki/Matching_wildcards) for matching wildcards.
//!
//! No escape characters are defined.
//!
//! - `?` matches exactly one occurrence of any character.
//! - `*` matches arbitrary many (including zero) occurrences of any character.
//!
//! Examples matching wildcards:
//! ``` rust
//! # extern crate wildmatch; use wildmatch::WildMatch;
//! assert!(WildMatch::new("cat").is_match("cat"));
//! assert!(WildMatch::new("*cat*").is_match("dog_cat_dog"));
//! assert!(WildMatch::new("c?t").is_match("cat"));
//! assert!(WildMatch::new("c?t").is_match("cot"));
//! ```
//! Examples not matching wildcards:
//! ``` rust
//! # extern crate wildmatch; use wildmatch::WildMatch;
//! assert!(!WildMatch::new("dog").is_match("cat"));
//! assert!(!WildMatch::new("*d").is_match("cat"));
//! assert!(!WildMatch::new("????").is_match("cat"));
//! assert!(!WildMatch::new("?").is_match("cat"));
//! ```

/// Wildcard matcher used to match strings.
type WildMatch struct {
	pattern []State
}

type State struct {
	NextChar    *rune
	HasWildcard bool
}

func (w *WildMatch) String() string {
	var sb strings.Builder
	for _, p := range w.pattern {
		sb.WriteString(string(*p.NextChar))
	}
	return sb.String()
}

// Constructor with pattern which can be used for matching.
func NewWildMatch(pattern string) *WildMatch {
	simplified := make([]State, 0)
	prevWasStar := false
	for _, currentChar := range pattern {
		copyCurrentChar := currentChar
		if currentChar == '*' {
			prevWasStar = true
		} else {
			s := State{
				NextChar:    &copyCurrentChar,
				HasWildcard: prevWasStar,
			}
			simplified = append(simplified, s)
			prevWasStar = false
		}
	}

	if len(pattern) > 0 {
		final := State{
			NextChar:    nil,
			HasWildcard: prevWasStar,
		}
		simplified = append(simplified, final)
	}

	return &WildMatch{
		pattern: simplified,
	}
}

// Indicates whether the matcher finds a match in the input string.
func (w *WildMatch) IsMatch(input string) bool {
	if len(w.pattern) == 0 {
		return false
	}

	patternIdx := 0
	for _, inputChar := range input {
		if patternIdx > len(w.pattern) {
			return false
		}

		p := w.pattern[patternIdx]

		if p.NextChar != nil && (*p.NextChar == '?' || *p.NextChar == inputChar) {
			patternIdx += 1
		} else if p.HasWildcard {
			if p.NextChar == nil {
				return true
			}
		} else {
			// Go back to last state with wildcard
			for {
				pattern := w.pattern[patternIdx]
				if pattern.HasWildcard {
					if pattern.NextChar != nil && (*pattern.NextChar == '?' || *pattern.NextChar == inputChar) {
						patternIdx += 1
					}
					break
				}
				if patternIdx == 0 {
					return false
				}
				patternIdx -= 1
			}
		}
	}
	return w.pattern[patternIdx].NextChar == nil
}

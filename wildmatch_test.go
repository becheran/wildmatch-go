package wildmatch_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/becheran/wildmatch-go/v2"
)

func TestIsMatchRandom(t *testing.T) {
	const permutations = 1_000

	randomString := func(length int) string {
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ*?"
		b := make([]byte, length)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		return string(b)
	}

	for patterIdx := 0; patterIdx < permutations; patterIdx++ {
		pattern := make([]byte, 100)
		for i := range pattern {
			pattern[i] = randomString(1)[0]
		}
		for i := 0; i < rand.Intn(25); i++ {
			pattern[rand.Intn(len(pattern))] = '*'
		}
		for i := 0; i < rand.Intn(25); i++ {
			pattern[rand.Intn(len(pattern))] = '?'
		}
		for inputIdx := 0; inputIdx < permutations; inputIdx++ {
			input := bytes.NewBuffer(nil)
			for _, p := range pattern {
				if p == '*' {
					input.WriteString(randomString(rand.Intn(15)))
				} else if p == '?' {
					input.WriteString(randomString(1))
				} else {
					input.WriteByte(p)
				}
			}

			m := wildmatch.NewWildMatch(string(pattern))
			if !m.Matches(input.String()) {
				t.Fatalf("Pattern\n%s\nInput:\n%s", m, input)
			}
		}
	}
}

func TestIsMatch(t *testing.T) {
	pattern := []struct {
		pattern string
		target  string
		result  bool
	}{
		// Match
		{pattern: "a*b", target: "a*xb", result: true},
		{pattern: "**", target: "cat", result: true},
		{pattern: "*", target: "cat", result: true},
		{pattern: "*?*", target: "cat", result: true},
		{pattern: "c*", target: "cat", result: true},
		{pattern: "c?*", target: "cat", result: true},
		{pattern: "???", target: "cat", result: true},
		{pattern: "c?t", target: "cat", result: true},
		{pattern: "cat", target: "cat", result: true},
		{pattern: "*cat", target: "cat", result: true},
		{pattern: "cat*", target: "cat", result: true},
		{pattern: "*cat*", target: "d&(*og_cat_dog", result: true},
		{pattern: "*?*", target: "d&(*og_cat_dog", result: true},
		{pattern: "*a*", target: "d&(*og_cat_dog", result: true},
		{pattern: "*", target: "*", result: true},
		{pattern: "*", target: "?", result: true},
		{pattern: "?", target: "?", result: true},
		{pattern: "wildcats", target: "wildcats", result: true},
		{pattern: "wild*cats", target: "wild?cats", result: true},
		{pattern: "wi*ca*s", target: "wildcats", result: true},
		{pattern: "wi*ca?s", target: "wildcats", result: true},
		{pattern: "*o?", target: "hog_cat_dog", result: true},
		{pattern: "*o?", target: "cat_dog", result: true},
		{pattern: "*at_dog", target: "cat_dog", result: true},
		{pattern: " ", target: " ", result: true},
		{pattern: "* ", target: "\n ", result: true},
		{pattern: "\n", target: "\n", result: true},
		{pattern: "*32", target: "432", result: true},
		{pattern: "*32", target: "332", result: true},
		{pattern: "*32", target: "3332", result: true},
		{pattern: "33*", target: "333", result: true},
		{pattern: " ", target: " ", result: true},

		// No Match
		{pattern: "*d*", target: "cat", result: false},
		{pattern: "*d", target: "cat", result: false},
		{pattern: "d*", target: "cat", result: false},
		{pattern: "*c", target: "cat", result: false},
		{pattern: "?", target: "cat", result: false},
		{pattern: "??", target: "cat", result: false},
		{pattern: "????", target: "cat", result: false},
		{pattern: "?????", target: "cat", result: false},
		{pattern: "*????", target: "cat", result: false},
		{pattern: "cats", target: "cat", result: false},
		{pattern: "cat?", target: "cat", result: false},
		{pattern: "cacat", target: "cat", result: false},
		{pattern: "cat*dog", target: "cat", result: false},
		{pattern: "cat?", target: "wildcats", result: false},
		{pattern: "cat*", target: "wildcats", result: false},
		{pattern: "*x*", target: "wildcats", result: false},
		{pattern: "*a", target: "wildcats", result: false},
		{pattern: "", target: "wildcats", result: false},
		{pattern: " ", target: "wildcats", result: false},
		{pattern: "???", target: "wildcats", result: false},
		{pattern: " ", target: "\n", result: false},
		{pattern: " ", target: "\t", result: false},
	}
	for _, p := range pattern {
		t.Run(p.pattern+"_"+p.target, func(t *testing.T) {
			m := wildmatch.NewWildMatch(p.pattern)
			result := m.Matches(p.target)
			if result != p.result {
				t.Fatalf("expected: %v, got: %v", p.result, result)
			}
		})
	}
}

func TestString(t *testing.T) {
	const input = "Lorem?ipsum*dolore*ea* ?????ata*."
	m := wildmatch.NewWildMatch(input)
	if m.String() != input {
		t.Fatalf("expected: %v, got: %v", input, m.String())
	}
}

func Benchmark_PatternCompile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		wildmatch.NewWildMatch("Lorem?ipsum*dolore*ea* ?????ata*.")
	}
}

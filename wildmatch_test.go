package wildmatch_test

import (
	"testing"

	"github.com/becheran/wildmatch-go"
)

type testCase struct {
	pattern string
	target  string
	result  bool
}

var pattern = []testCase{
	// Match
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

func TestIsMatch(t *testing.T) {
	for _, p := range pattern {
		t.Run(p.pattern+"_"+p.target, func(t *testing.T) {
			m := wildmatch.NewWildMatch(p.pattern)
			result := m.IsMatch(p.target)
			if result != p.result {
				t.Fatalf("expected: %v, got: %v", p.result, result)
			}
		})
	}
}

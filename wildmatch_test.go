package wildmatch

import (
	"testing"
)

type testCase struct {
	pattern string
	target  string
	result  bool
}

var pattern = []testCase{
	// Simple match
	//testCase{pattern: "**", target: "cat", result: true},
	//testCase{pattern: "*", target: "cat", result: true},
	testCase{pattern: "*?*", target: "cat", result: true},
	testCase{pattern: "c*", target: "cat", result: true},
	testCase{pattern: "c?*", target: "cat", result: true},
	testCase{pattern: "???", target: "cat", result: true},
	testCase{pattern: "c?t", target: "cat", result: true},
	testCase{pattern: "cat", target: "cat", result: true},
	testCase{pattern: "*cat", target: "cat", result: true},
	testCase{pattern: "cat*", target: "cat", result: true},
	// Simple No Match
	testCase{pattern: "*d*", target: "cat", result: false},
	testCase{pattern: "*d", target: "cat", result: false},
	testCase{pattern: "d*", target: "cat", result: false},
	testCase{pattern: "*c", target: "cat", result: false},
	testCase{pattern: "?", target: "cat", result: false},
	testCase{pattern: "??", target: "cat", result: false},
	testCase{pattern: "????", target: "cat", result: false},
	testCase{pattern: "?????", target: "cat", result: false},
	testCase{pattern: "*????", target: "cat", result: false},
	testCase{pattern: "cats", target: "cat", result: false},
	testCase{pattern: "cat?", target: "cat", result: false},
	testCase{pattern: "cacat", target: "cat", result: false},
	testCase{pattern: "cat*dog", target: "cat", result: false},
}

func TestIsMatch(t *testing.T) {
	for _, p := range pattern {
		t.Run(p.pattern+"_"+p.target, func(t *testing.T) {
			m := NewWildMatch(p.pattern)
			result := m.IsMatch(p.target)
			if result != p.result {
				t.Fatalf("expected: %v, got: %v", p.result, result)
			}
		})
	}
}

/*

#[test_case("**")]
#[test_case("*")]
#[test_case("*?*")]
#[test_case("c*")]
#[test_case("c?*")]
#[test_case("???")]
#[test_case("c?t")]
#[test_case("cat")]
#[test_case("*cat")]
#[test_case("cat*")]
fn is_match(pattern: &str) {
	let m = WildMatch::new(pattern);
	assert!(m.is_match("cat"));
}

#[test_case("*d*")]
#[test_case("*d")]
#[test_case("d*")]
#[test_case("*c")]
#[test_case("?")]
#[test_case("??")]
#[test_case("????")]
#[test_case("?????")]
#[test_case("*????")]
#[test_case("cats")]
#[test_case("cat?")]
#[test_case("cacat")]
#[test_case("cat*dog")]
fn no_match(pattern: &str) {
	let m = WildMatch::new(pattern);
	assert_false!(m.is_match("cat"));
}

#[test_case("cat?", "wildcats")]
#[test_case("cat*", "wildcats")]
#[test_case("*x*", "wildcats")]
#[test_case("*a", "wildcats")]
#[test_case("", "wildcats")]
#[test_case(" ", "wildcats")]
#[test_case("???", "wildcats")]
fn no_match_long(pattern: &str, expected: &str) {
	let m = WildMatch::new(pattern);
	assert_false!(m.is_match(expected))
}

#[test_case("*cat*", "d&(*og_cat_dog")]
#[test_case("*?*", "d&(*og_cat_dog")]
#[test_case("*a*", "d&(*og_cat_dog")]
#[test_case("*", "*")]
#[test_case("*", "?")]
#[test_case("?", "?")]
#[test_case("wildcats", "wildcats")]
#[test_case("wild*cats", "wild?cats")]
#[test_case("wi*ca*s", "wildcats")]
#[test_case("wi*ca?s", "wildcats")]
#[test_case("*o?", "hog_cat_dog")]
#[test_case("*o?", "cat_dog")]
#[test_case("*at_dog", "cat_dog")]
#[test_case(" ", " ")]
#[test_case("* ", "\n ")]
#[test_case("\n", "\n", name = "special_chars")]
#[test_case("*32", "432")]
#[test_case("*32", "332")]
#[test_case("33*", "333")]
fn match_long(pattern: &str, expected: &str) {
	let m = WildMatch::new(pattern);
	assert!(m.is_match(expected))
}

#[test]
fn print_string() {
	let m = WildMatch::new("Foo/Bar");
	assert_eq!("Foo/Bar", m.to_string());
}

#[test]
fn to_string_f() {
	let m = WildMatch::new("F");
	assert_eq!("F", m.to_string());
}

#[test]
fn to_string_empty() {
	let m = WildMatch::new("");
	assert_eq!("", m.to_string());
}
*/

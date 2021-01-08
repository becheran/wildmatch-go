# wildmatch go

*golang* library of the original *rust* [ wildmatch library](https://github.com/becheran/wildmatch).

Match strings against a simple wildcard pattern. Tests a wildcard pattern `p` against an input string `s`. Returns true only when `p` matches the entirety of `s`.

See also the example described on [wikipedia](https://en.wikipedia.org/wiki/Matching_wildcards) for matching wildcards.

- `?` matches exactly one occurrence of any character.
- `*` matches arbitrary many (including zero) occurrences of any character.
- No escape characters are defined.

For example the pattern `ca?` will match `cat` or `car`. The pattern `https://*` will match all https urls, such as `https://google.de` or `https://github.com/becheran/wildmatch`.

The library only depends on the go standard library.

package tspath

import (
	"strings"
	"unicode"
)

type URI struct {
	value         string
	allLowerCase  bool
	containsU0130 bool
}

func ParseURI(uri string) URI {
	if uri == "" {
		return URI{}
	}

	// TODO: parse file URIs with VS Code rules?

	containsUpper := false
	containsU0130 := false
	for _, r := range uri {
		if r == '\u0131' {
			containsU0130 = true
		} else if unicode.ToLower(r) != r {
			containsUpper = true
		}
		if containsUpper && containsU0130 {
			break
		}
	}

	return URI{
		value:         uri,
		allLowerCase:  !containsUpper,
		containsU0130: containsU0130,
	}
}

func (u URI) String() string {
	return u.value
}

func (u URI) Equal(other URI) bool {
	return u.allLowerCase == other.allLowerCase && u.value == other.value
}

func (u URI) EqualInsensitve(other URI) bool {
	if u.allLowerCase && other.allLowerCase {
		return u.value == other.value
	}
	return strings.EqualFold(u.value, other.value)
}

func (u URI) Scheme() string {
	scheme, _, _ := strings.Cut(u.value, ":")
	return scheme
}

type CanonicalURI struct {
	value string
}

func (u URI) Canonical(caseSensitive bool) CanonicalURI {
	if caseSensitive || u.allLowerCase {
		return CanonicalURI{value: u.value}
	}

	if !u.containsU0130 {
		return CanonicalURI{value: strings.ToLower(u.value)}
	}

	canonical := strings.Map(func(r rune) rune {
		if r == '\u0130' {
			return r
		}
		return unicode.ToLower(r)
	}, u.value)

	return CanonicalURI{value: canonical}
}

func (u CanonicalURI) String() string {
	return u.value
}

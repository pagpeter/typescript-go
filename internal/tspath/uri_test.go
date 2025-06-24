package tspath

import "testing"

func TestParseURI(t *testing.T) {
	tests := []struct {
		uri      string
		expected tspath.URI
	}{
		{"file:///path/to/file.ts", tspath.Parse("file:///path/to/file.ts")},
		{"FILE:///path/to/file.ts", tspath.Parse("file:///path/to/file.ts")},
		{"file:///path/to/FILE.ts", tspath.Parse("file:///path/to/file.ts")},
		{"file:///path/to/FILE\u0130.ts", tspath.Parse("file:///path/to/file\u0130.ts")},
	}

	for _, test := range tests {
		t.Run(test.uri, func(t *testing.T) {
			result := tspath.Parse(test.uri)
			if !result.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

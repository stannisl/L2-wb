package main

import (
	"testing"
)

func TestUnpackStr(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"qwe\\4\\5", "qwe45", false},
		{"a0b", "b", false},
		{"3abc", "", true},
		{"qwe\\", "qwe", false},
		{"", "", false},
		{"45", "", true},
		{"qwe\\45", "qwe44444", false},
	}

	for _, tt := range tests {
		got, err := unpackStr(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("unpackStr(%q) error = %v, wantErr = %v", tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.expected {
			t.Errorf("unpackStr(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

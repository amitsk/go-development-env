package models

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Hero is a single record the API exposes.
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ValidateName reports whether a hero name is acceptable.
func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if utf8.RuneCountInString(name) < 2 {
		return fmt.Errorf("name is too short")
	}
	return nil
}

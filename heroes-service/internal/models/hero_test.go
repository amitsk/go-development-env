package models

import "testing"

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "valid", input: "Superman", wantErr: false},
		{name: "empty", input: "", wantErr: true},
		{name: "spaces only", input: "   ", wantErr: true},
		{name: "too short", input: "A", wantErr: true},
		{name: "unicode two runes", input: "闪闪", wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateName(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ValidateName(%q) err=%v wantErr=%v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func FuzzValidateName(f *testing.F) {
	f.Add("Superman")
	f.Add("")
	f.Add("闪")
	f.Fuzz(func(t *testing.T, name string) {
		_ = ValidateName(name)
	})
}

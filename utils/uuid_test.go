package utils

import (
	"strconv"
	"testing"
)

func TestGetUniqueID(t *testing.T) {
	tests := []struct {
		name string
		want func(string) bool
	}{
		{"length 32", func(id string) bool { return len(id) == 32 }},
		{"base16 encoded", func(id string) bool {
			for i := 0; i < 32; i += 8 {
				if _, err := strconv.ParseInt(string(id[i:i+8]), 16, 64); err != nil {
					print(err.Error())
					return false
				}
			}
			return true
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUniqueID(); !tt.want(got) {
				t.Errorf("GetUniqueID() = %v failed expectation", got)
			}
		})
	}
}

package chart

import "testing"

func TestCreateBar(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TEST"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateBar()
		})
	}
}

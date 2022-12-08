package chart

import "testing"

func TestLineExamples_Examples(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			li := LineExamples{}
			li.Examples()
		})
	}
}

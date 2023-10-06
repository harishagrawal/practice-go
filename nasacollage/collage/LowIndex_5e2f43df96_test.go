// Test generated by RoostGPT for test practice-go-nasacollage using AI Type Azure Open AI and AI Model roost-gpt4-32k

package collage

import (
	"testing"
)

type Bar struct {
	H int
}

type BarGraph []Bar

func TestLowIndex(t *testing.T) {
	tests := []struct {
		name  string
		bars  BarGraph
		want  int
	}{
		{
			name:  "No Bars",
			bars:  BarGraph{},
			want:  0,
		},
		{
			name:  "Valid Bars",
			bars:  BarGraph{{H: 5}, {H: 1}, {H: 3}},
			want:  1,
		},
		{
			name:  "Duplicate Low Bars",
			bars:  BarGraph{{H: 5}, {H: 1}, {H: 1}},
			want:  1,
		},
		{
			name:  "Bars with Negative Height",
			bars:  BarGraph{{H: 5}, {H: -1}, {H: 3}},
			want:  1,
		},
		{
			name:  "Single Bar Test",
			bars:  BarGraph{{H: 5}},
			want:  0,
		},
		{
			name:  "Test with Equal Bar Heights",
			bars:  BarGraph{{H: 1}, {H: 1}, {H: 1}},
			want:  0,
		},
		{
			name:  "Test Bar Heights with Floats (returns int portion of min)",
			bars:  []Bar{{H: 1.5}, {H: 2.1}, {H: 1.9}},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bars.LowIndex(); got != tt.want {
				t.Errorf("LowIndex() = %v, want %v", got, tt.want)
			} else {
				t.Log("Success: Expected outcome matched found outcome")
			}
		})
	}
}

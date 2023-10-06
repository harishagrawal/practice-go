// Test generated by RoostGPT for test practice-go-warriors using AI Type Open AI and AI Model gpt-4


/*
1. Scenario: Test with an empty map
   Given an empty map, the function should return an empty slice.

2. Scenario: Test with a map of one key-value pair
   Given a map with one key-value pair, the function should return a slice with one key.

3. Scenario: Test with a map of multiple key-value pairs
   Given a map with multiple key-value pairs, the function should return a slice containing all keys in the map.

4. Scenario: Test with a map with duplicate keys
   Given a map with duplicate keys, the function should return a slice containing only unique keys.

5. Scenario: Test with a map with keys of varying lengths
   Given a map with keys of varying lengths, the function should return a slice containing all keys, preserving their respective lengths.

6. Scenario: Test with a map with keys containing special characters
   Given a map with keys containing special characters, the function should return a slice containing all keys, including the special characters.

7. Scenario: Test with a map with keys containing whitespace
   Given a map with keys containing whitespace, the function should return a slice containing all keys, including the whitespace.

8. Scenario: Test with a map containing both positive and negative values
   Given a map containing both positive and negative values, the function should return a slice containing all keys, regardless of the value associated with each key.

9. Scenario: Test with a large map
   Given a large map, the function should return a slice containing all keys without causing any memory issues.

10. Scenario: Test with a map containing keys of different data types
   Given a map containing keys of different data types, the function should return a slice containing all string keys and ignore keys of other data types.
*/
package functionfrequency_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yourusername/functionfrequency"
)

func TestKeysToSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]int
		expected []string
	}{
		{
			name:     "Test with an empty map",
			input:    map[string]int{},
			expected: []string{},
		},
		{
			name:     "Test with a map of one key-value pair",
			input:    map[string]int{"key1": 1},
			expected: []string{"key1"},
		},
		{
			name:     "Test with a map of multiple key-value pairs",
			input:    map[string]int{"key1": 1, "key2": 2, "key3": 3},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name:     "Test with a map with keys of varying lengths",
			input:    map[string]int{"k": 1, "ke": 2, "key": 3},
			expected: []string{"k", "ke", "key"},
		},
		{
			name:     "Test with a map with keys containing special characters",
			input:    map[string]int{"k$y": 1, "k@y": 2, "k#y": 3},
			expected: []string{"k$y", "k@y", "k#y"},
		},
		{
			name:     "Test with a map with keys containing whitespace",
			input:    map[string]int{"k y": 1, "k  y": 2, "k   y": 3},
			expected: []string{"k y", "k  y", "k   y"},
		},
		{
			name:     "Test with a map containing both positive and negative values",
			input:    map[string]int{"key1": 1, "key2": -2, "key3": 3},
			expected: []string{"key1", "key2", "key3"},
		},
		{
			name:     "Test with a large map",
			input:    getLargeMap(),
			expected: getLargeSlice(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functionfrequency.KeysToSlice(tt.input)
			sort.Strings(result)
			sort.Strings(tt.expected)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func getLargeMap() map[string]int {
	m := make(map[string]int)
	for i := 0; i < 10000; i++ {
		m["key"+string(i)] = i
	}
	return m
}

func getLargeSlice() []string {
	s := make([]string, 0, 10000)
	for i := 0; i < 10000; i++ {
		s = append(s, "key"+string(i))
	}
	return s
}

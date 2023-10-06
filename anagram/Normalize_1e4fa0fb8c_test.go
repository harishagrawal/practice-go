// Test generated by RoostGPT for test practice-go-warriors using AI Type Open AI and AI Model gpt-4


/*
1. Test with a standard string: Check how the function handles a string containing both uppercase and lowercase letters. E.g. "HelloWorld". The expected output should be "dehllloorw".

2. Test with a string containing special characters: The function should ignore special characters and only return the sorted, lowercase letters. E.g. "Hello, World!". The expected output should be "dehllloorw".

3. Test with a string containing numbers: Check how the function handles a string containing numbers. E.g. "Hello123". The expected output should be "ehllo".

4. Test with an empty string: Check how the function handles an empty string. The expected output should be an empty string.

5. Test with a string containing only uppercase letters: The function should convert all letters to lowercase and sort them. E.g. "HELLO". The expected output should be "ehllo".

6. Test with a string containing only lowercase letters: The function should just sort the letters. E.g. "hello". The expected output should be "ehllo".

7. Test with a string containing only special characters: The function should return an empty string as there are no letters to sort. E.g. "!!!@@@###". The expected output should be an empty string.

8. Test with a string containing only numbers: The function should return an empty string as there are no letters to sort. E.g. "1234567890". The expected output should be an empty string.

9. Test with a string containing spaces: The function should ignore spaces and only return the sorted, lowercase letters. E.g. "Hello World". The expected output should be "dehllloorw".

10. Test with a string containing non-English letters: Check how the function handles non-English letters. E.g. "Héllo". The expected output should be "ellho".
*/
package anagram

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"standard string", "HelloWorld", "dehllloorw"},
		{"string with special characters", "Hello, World!", "dehllloorw"},
		{"string with numbers", "Hello123", "ehllo"},
		{"empty string", "", ""},
		{"uppercase string", "HELLO", "ehllo"},
		{"lowercase string", "hello", "ehllo"},
		{"string with special characters only", "!!!@@@###", ""},
		{"string with numbers only", "1234567890", ""},
		{"string with spaces", "Hello World", "dehllloorw"},
		{"non-English letters", "Héllo", "ellho"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Normalize(tc.input)
			if result != tc.expected {
				t.Errorf("Normalize(%q) = %q; expected %q", tc.input, result, tc.expected)
			} else {
				t.Logf("Success: Normalize(%q) = %q", tc.input, result)
			}
		})
	}
}

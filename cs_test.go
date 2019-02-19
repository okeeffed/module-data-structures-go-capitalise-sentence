package caps

import "testing"

func TestCapitaliseSentence(t *testing.T) {
	for _, tt := range testCases {
		res := capitaliseSentence(tt.input)

		if res != tt.expected {
			t.Errorf("FAIL: Expected %s but got %s", tt.expected, res)
			return
		}

		t.Logf("PASS: Expected %s and got %s", tt.expected, res)
	}
}

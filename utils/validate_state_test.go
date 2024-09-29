package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateValidation(t *testing.T) {
	t.Parallel()

	var cases = []struct {
		state  string
		expect bool
	}{
		{"invalid", false},
		{"open", true},
		{"closed", true},
		{"accepted", true},
		{"investigating", true},
	}

	for _, testcase := range cases {
		actual := IsValidState(testcase.state)
		assert.Equal(t, testcase.expect, actual, "Incorrect result from state validating!")
	}
}

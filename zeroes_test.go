package clearmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestZeroes(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name: "flat",
			input: map[string]interface{}{
				"empty_int":        0,
				"non_empty_int":    1,
				"empty_string":     "",
				"non_empty_string": "foo",
				"false_bool":       false,
				"true_bool":        true,
			},
			output: map[string]interface{}{
				"non_empty_int":    1,
				"non_empty_string": "foo",
				"true_bool":        true,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.output, Zeroes(test.input)); diff != "" {
				t.Errorf("(- want, + got):\n %s", diff)
			}
		})
	}
}

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
		{
			name: "one-layer-deep",
			input: map[string]interface{}{
				"recursive": map[string]interface{}{
					"empty_int":        0,
					"non_empty_int":    1,
					"empty_string":     "",
					"non_empty_string": "foo",
					"false_bool":       false,
					"true_bool":        true,
				},
				"recursive_array": []map[string]interface{}{
					{
						"empty_int":        0,
						"non_empty_int":    1,
						"empty_string":     "",
						"non_empty_string": "foo",
						"false_bool":       false,
						"true_bool":        true,
					},
				},
				"empty_int":        0,
				"non_empty_int":    1,
				"empty_string":     "",
				"non_empty_string": "foo",
				"false_bool":       false,
				"true_bool":        true,
			},
			output: map[string]interface{}{
				"recursive": map[string]interface{}{
					"non_empty_int":    1,
					"non_empty_string": "foo",
					"true_bool":        true,
				},
				"recursive_array": []map[string]interface{}{
					{
						"non_empty_int":    1,
						"non_empty_string": "foo",
						"true_bool":        true,
					},
				},
				"non_empty_int":    1,
				"non_empty_string": "foo",
				"true_bool":        true,
			},
		},
		{
			name: "empty-recursive",
			input: map[string]interface{}{
				"recursive": map[string]interface{}{
					"empty_int":    0,
					"empty_string": "",
					"false_bool":   false,
				},
				"recursive_array": []map[string]interface{}{
					{
						"empty_int":    0,
						"empty_string": "",
						"false_bool":   false,
					},
				},
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
		{
			name: "empty",
			input: map[string]interface{}{
				"recursive": map[string]interface{}{
					"empty_int":    0,
					"empty_string": "",
					"false_bool":   false,
				},
				"recursive_array": []map[string]interface{}{
					{
						"empty_int":    0,
						"empty_string": "",
						"false_bool":   false,
					},
				},
				"empty_int":    0,
				"empty_string": "",
				"false_bool":   false,
			},
			output: nil,
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

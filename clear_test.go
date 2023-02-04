package clearmap

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClear(t *testing.T) {
	tests := []struct {
		name   string
		input  map[string]interface{}
		output map[string]interface{}
	}{
		{
			name: "flat",
			input: map[string]interface{}{
				"empty_byte":        byte(0),
				"non_empty_byte":    byte(1),
				"empty_complex":     complex(0, 0),
				"non_empty_complex": complex(2, 4),
				"empty_rune":        rune(0),
				"non_empty_rune":    'i',
				"empty_int":         0,
				"non_empty_int":     1,
				"empty_float":       0.0,
				"non_empty_float":   1.0,
				"zero_float":        float64(0),
				"empty_string":      "",
				"non_empty_string":  "foo",
				"false_bool":        false,
				"true_bool":         true,
			},
			output: map[string]interface{}{
				"non_empty_rune":    'i',
				"non_empty_complex": complex(2, 4),
				"non_empty_byte":    byte(1),
				"non_empty_int":     1,
				"non_empty_float":   1.0,
				"non_empty_string":  "foo",
				"true_bool":         true,
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
		{
			name: "empties",
			input: map[string]interface{}{
				"empty_map":   map[string]interface{}{},
				"empty_array": []map[string]interface{}{},
				"nil":         nil,
			},
			output: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.output, Clear(test.input)); diff != "" {
				t.Errorf("(- want, + got):\n %s", diff)
			}
		})
	}
}

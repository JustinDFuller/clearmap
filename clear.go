// package clearmap implements a utility for clearing maps of zero, empty, and nil values.
package clearmap

// Clear clears maps of all zero, empty, and nil values.
// If the resulting map is empty, nil is returned.
//
// # Supported Types
//
// - Ints (int64, int32, int16, int8, int)
// - Uints (uint64, uint32, uint16, uint8, uint)
// - Strings (string, rune, byte)
// - Booleans
// - Iterable types (maps, arrays)
//
// It will work recursively on maps and slice types.
//
// Changes are made in-place but the map is returned for convenience.
func Clear(m map[string]interface{}) map[string]interface{} {
	for key, val := range m {
		if val == nil {
			delete(m, key)
			continue
		}

		switch v := val.(type) {
		case complex64, complex128:
			if v == complex(0, 0) {
				delete(m, key)
			}
		case uint8, uint16, uint32, uint64, uint, uintptr:
			if v == 0 || v == byte(0) {
				delete(m, key)
			}
		case int64, int32, int16, int8, int:
			if v == 0 || v == rune(0) || v == complex(0, 0) {
				delete(m, key)
			}
		case float32, float64:
			if v == 0.0 {
				delete(m, key)
			}
		case string:
			if v == "" {
				delete(m, key)
			}
		case bool:
			if v == false {
				delete(m, key)
			}
		case map[string]interface{}:
			Clear(v)
			if len(v) == 0 {
				delete(m, key)
			}
		case []map[string]interface{}:
			for i := len(v) - 1; i >= 0; i-- {
				Clear(v[i])
				if len(v[i]) == 0 {
					v = v[:i]
				}
			}
			if len(v) == 0 {
				delete(m, key)
			}
		}
	}

	if len(m) == 0 {
		return nil
	}

	return m
}

package clearmap

func Zeroes(m map[string]interface{}) map[string]interface{} {
	for key, val := range m {
		switch v := val.(type) {
		case int64, int32, int:
			if v == 0 {
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
			Zeroes(v)
			if len(v) == 0 {
				delete(m, key)
			}
		case []map[string]interface{}:
			for i := len(v) - 1; i >= 0; i-- {
				Zeroes(v[i])
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

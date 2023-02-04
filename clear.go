package clearmap

func Clear(m map[string]interface{}) map[string]interface{} {
	for key, val := range m {
		if val == nil {
			delete(m, key)
			continue
		}

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

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
		}
	}

	return m
}

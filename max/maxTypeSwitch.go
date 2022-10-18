package max

func maxTypeSwitch(a, b interface{}) (bool, interface{}) {
	switch a.(type) {
	case int:
		if bi, ok := b.(int); ok {
			if a.(int) > bi {
				return true, a
			}
			return true, bi
		}
	case float64:
		if bi, ok := b.(float64); ok {
			if a.(float64) > bi {
				return true, a
			}
			return true, bi
		}
	case string:
		if bi, ok := b.(string); ok {
			if a.(string) > bi {
				return true, a
			}
			return true, bi
		}
	}

	return false, nil
}

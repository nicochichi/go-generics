package add

func addTypeSwitch(a, b interface{}) (bool, interface{}) {
	switch a.(type) {
	case int:
		if bi, ok := b.(int); ok {
			return true, a.(int) + bi
		}
	case float64:
		if bi, ok := b.(float64); ok {
			return true, a.(float64) + bi
		}
	case string:
		if bi, ok := b.(string); ok {
			return true, a.(string) + bi
		}
	}

	return false, nil
}

package max

import "reflect"

func maxReflection(a, b interface{}) (bool, interface{}) {
	if reflect.TypeOf(a).Kind() == reflect.Int && reflect.TypeOf(b).Kind() == reflect.Int {
		if a.(int) > b.(int) {
			return true, a
		}
		return true, b
	}
	if reflect.TypeOf(a).Kind() == reflect.Float64 && reflect.TypeOf(b).Kind() == reflect.Float64 {
		if a.(float64) > b.(float64) {
			return true, a
		}
		return true, b
	}
	if reflect.TypeOf(a).Kind() == reflect.String && reflect.TypeOf(b).Kind() == reflect.String {
		if a.(string) > b.(string) {
			return true, a
		}
		return true, b
	}
	return false, nil
}

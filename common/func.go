package common

// return true if data string is empty
func IsEmptyString(data string) bool {
	if data == "" {
		return true
	}
	return false
}

//func IsNil(v interface{}) bool {
//	vi := reflect.TypeOf(v)
//	if vi.Kind() == reflect.Ptr {
//		return vi.Comparable()
//	}
//	return false
//}

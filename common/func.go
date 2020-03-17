package common

import (
	"time"
)

const (
	// CST format time layout
	CENTRAL_STANDARD_TIME_LAYOUT = "2006-01-02 15:04:05 +0800 CST"
)

// return true if data string is empty
func IsEmptyString(data string) bool {
	if data == "" {
		return true
	}
	return false
}

// return err if format time has failed
func FormatTime(layout, timeValue string) (time.Time, error) {

	return time.Parse(layout, timeValue)
}

//func IsNil(v interface{}) bool {
//	vi := reflect.TypeOf(v)
//	if vi.Kind() == reflect.Ptr {
//		return vi.Comparable()
//	}
//	return false
//}

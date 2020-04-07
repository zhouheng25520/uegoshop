package common

import (
	"fmt"
	"reflect"
	"time"
)

const (
	// CST format time layout
	CentralStandardTimeLayout = "2006-01-02 15:04:05 +0800 CST"
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


// return true, if value eq nil
func IsNil(any interface{}) bool {
	re := false
	if any != nil {
		v := reflect.ValueOf(any)

		if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
			re = v.IsNil()
			if !re {
				for {
					fmt.Println(v.Type())
					v2 := v.Elem()
					if v2.Kind() != reflect.Ptr && v2.Kind() != reflect.Interface {
						break
					}
					re = v2.IsNil()
					if re {
						break
					}
					v = v2
				}
			}

		}
	}
	return re
}
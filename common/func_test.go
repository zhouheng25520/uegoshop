package common

import (
	"reflect"
	"testing"
)

func TestIsEmptyString(t *testing.T) {
	//mapTest := map[string]string{}
	//mapTest["key"] = "keyValue"
	tests := []struct{
		//input interface{}
		input string
		output bool
	}{
		{"", true},
		{"test", false},
		{"32132", false},
		//{0, true},
		//{123, false},
		//{make(map[string]string), true},
		//{mapTest, false},
	}
	for i, test := range tests {
		if isEmpty := IsEmptyString(test.input); isEmpty != test.output {
			t.Errorf("#%d data type[%v], should be %v, but got %v",
				i, reflect.TypeOf(test.input), test.output, isEmpty)
		}
	}
}

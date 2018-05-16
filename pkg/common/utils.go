package common

import (
	"reflect"
	"github.com/google/logger"
)

/*
	GetFromSlice() takes slice and searches an item in it
	returns number in slice and item in slice;
			-1 and nil if found nothing
*/
func GetFromSlice(items []interface{}, item interface{}) (int, interface{}) {
	for i := range items {
		if reflect.DeepEqual(items[i], item) {
			return i, items[i]
		}
	}
	return -1, nil
}

func GetFromSliceF(items []interface{}, item interface{}, fn func(interface{}) interface{}) (int, interface{}) {
	for i := range items {
		if reflect.DeepEqual(fn(items[i]), item) {
			return i, items[i]
		}
	}
	return -1, nil
}


/*
	ToInterface() turns slice []T into []interface{}
 */
func ToInterface(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		logger.Error("There wasn't a slice sent as an argument to InterfaceSlice")
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i:=0; i<s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

/*
	IsError() tells if the interface{} value contains error
*/
func IsError(item interface{}) bool {
	if _, ok := item.(error); !ok {
		if _, ok := item.(*error); !ok {
			return false
		}
	}
	return true
}
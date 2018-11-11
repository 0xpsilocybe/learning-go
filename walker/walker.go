package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	value := getValue(x)
	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}
	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walkValue(value.Field(i))
		}
	case reflect.Array,
         reflect.Slice:
		 for i := 0; i < value.Len(); i++ {
			 walkValue(value.Index(i))
		 }
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walkValue(value.MapIndex(key))
		}
	}
}

func getValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)
	if value.Kind() ==reflect.Ptr {
		value = value.Elem()
	}
	return value
}


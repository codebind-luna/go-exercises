package main

import (
	"log"
	"reflect"
)

func flatten(input interface{}) []int {
	var result []int

	val := reflect.ValueOf(input)

	switch val.Kind() {
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i).Interface()

			if reflect.TypeOf(elem).Kind() == reflect.Slice {
				result = append(result, flatten(elem)...)
			} else {
				v, ok := elem.(int)

				if ok {
					result = append(result, v)
				}
			}
		}
	}

	return result
}

func main() {
	nestedArray := []interface{}{
		1, 2, []interface{}{3, 4}, []interface{}{5, []interface{}{6, 7}, 8},
		9, []interface{}{10, []interface{}{11, 12}},
	}

	log.Println(nestedArray)
	flattened := flatten(nestedArray)
	log.Println(flattened)
}

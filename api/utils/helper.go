package helper

import (
	"fmt"
	"reflect"
)

func ShowStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	fmt.Printf("Struct: %s\n", typ.Name())
	fmt.Println("Fields:")

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)

		fmt.Printf("- %s (%s): %v\n", field.Name, field.Type, value)
	}
}

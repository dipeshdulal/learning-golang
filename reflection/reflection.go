package reflection

import (
	"fmt"
	"reflect"
)

//GetDataType Get the data type
func GetDataType(d interface{}) {
	varType := reflect.TypeOf(d)
	fmt.Println("Data type is: ", varType)
}

//GetFields Gets the fields and prints
func GetFields(d interface{}) {
	varType := reflect.TypeOf(d)
	fieldZero := varType.Field(0)

	fmt.Println("Data type is: ", varType)
	fmt.Println("fieldZero is: ", fieldZero)
}

//ChangeValueWithReflection Change the value of variable using reflection
func ChangeValueWithReflection(d interface{}) {
	field := reflect.ValueOf(d).Elem().Field(0)

	// valueOf gives reflection pointer, elem dereferences pointer, field gets value from pointer

	str := "Ola Comoestas"
	field.SetString(str)

}

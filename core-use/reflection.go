package core_use

import (
	"fmt"
	"reflect"
)

type Foo struct{}

func (f Foo) Hello() string {
	return fmt.Sprintf("hello world")
}

type S struct {
	a int
	b float64
}

type Example struct {
	A int
	B string
	C bool
}

func emptyFields(s interface{}, fieldNames []string) []string {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体类型")
		return nil
	}

	t := v.Type()
	emptyFields := []string{}

	for _, fieldName := range fieldNames {
		index, ok := t.FieldByName(fieldName)
		if !ok {
			fmt.Printf("结构体中没有名为 %s 的字段\n", fieldName)
			continue
		}
		field := v.Field(index.Index[0])
		zeroValue := reflect.Zero(field.Type())

		if field.Interface() == zeroValue.Interface() {
			emptyFields = append(emptyFields, fieldName)
		}
	}

	return emptyFields
}

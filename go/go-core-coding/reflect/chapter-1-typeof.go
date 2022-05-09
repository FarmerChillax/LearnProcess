package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string "学生姓名"
	Age  int    `a:"111"b:"3333"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)

	fieldName, ok := rt.FieldByName("Name")

	// 取 tag 数据
	if ok {
		fmt.Println(fieldName.Tag)
	}
	fieldAge, ok2 := rt.FieldByName("Age")

	// 可以像 JSON 一样，取 tag 里的数据，多个 tag 直接无逗号
	// tag 不需要引号
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a"))
		fmt.Println(fieldAge.Tag.Get("b"))
	}

	fmt.Println("type_name:", rt.Name())
	fmt.Println("type_NumField:", rt.NumField())
	fmt.Println("type_PkgPath:", rt.PkgPath())
	fmt.Println("type_String:", rt.String())

	fmt.Println("type_Kind.String:", rt.Kind().String())
	fmt.Println("type_String()=", rt.String())

	// 获取结构类型的字段名称
	for i := 0; i < rt.NumField(); i++ {
		fmt.Printf("type.Field[%d].Name := %v \n", i, rt.Field(i).Name)
	}

	sc := make([]int, 10)
	sc = append(sc, 1, 2, 3)
	sct := reflect.TypeOf(sc)

	// 获取 slice 元素的 Type
	scet := sct.Elem()

	fmt.Println("slice element type.Kind()=", scet.Kind())
	fmt.Printf("slice element type.Kind()=%d\n", scet.Kind())
	fmt.Println("slice element type.String()=", scet.String())

	fmt.Println("slice element type.Name()=", scet.Name())
	fmt.Println("slice element type.NumMethod()=", scet.NumMethod())
	fmt.Println("slice element type.PkgPath()=", scet.PkgPath())
	fmt.Println("slice element type.PkgPath()=", sct.PkgPath())

}

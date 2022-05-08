package main

import (
	"fmt"
	"reflect"
)

type Name string

type Age int32

type Course []string

type Students []Name

// Alias
type Username = string

type Password = Username

func main() {
	nameType := Name("farmer")
	fmt.Printf("%v -> %T\n", nameType, nameType)

	nameAlias := Username("farmer")
	fmt.Printf("%v -> %T\n", nameAlias, nameAlias)
	fmt.Println(reflect.TypeOf(nameAlias), reflect.TypeOf(nameType))
}

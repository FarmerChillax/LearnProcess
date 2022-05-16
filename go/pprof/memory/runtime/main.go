package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	refMemStats := reflect.TypeOf(*memStats)
	fmt.Println(refMemStats.Kind())
	// for _, elem := range  {
	// 	fmt.Println(elem)
	// }
}

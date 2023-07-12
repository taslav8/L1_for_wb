package main

import (
	"fmt"
	"reflect"
)

func detectorOne(in interface{}) {
	fmt.Println(reflect.TypeOf(in))
}
func detectorTwo(in interface{}) {
	fmt.Println(reflect.ValueOf(in).Kind())
}
func detectorThree(in interface{}) {
	fmt.Println(fmt.Sprintf("%T", in))
}

func main() {
	var a int64
	detectorOne(a)
	detectorTwo(a)
	detectorThree(a)
}

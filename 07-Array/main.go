package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var array [5]int
	array[0] = 1
	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))
	fmt.Println(unsafe.Sizeof(array))
	fmt.Println(len(array))

	array2 := [5][4]int{}
	for a,i :=range array2{
		for b,_:=range i{
			array2[a][b]=a+b
		}
	}

	fmt.Println(array2)
}
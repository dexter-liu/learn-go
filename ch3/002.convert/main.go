package main

import (
	"fmt"
)

func main() {
	convertType()
	fmt.Println("finished")
}

func convertType() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("pannic啦", r)
			// debug.PrintStack()
		}
	}()

	var a interface{} = "string aaa"

	b := a.(int)
	fmt.Println(b)
}

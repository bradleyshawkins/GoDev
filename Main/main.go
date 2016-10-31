package main

import (
	"fmt"

	"github.com/bradleyshawkins/GoDev/MyStruct"
	"github.com/bradleyshawkins/stringutil"
)

func main() {
	fmt.Println(stringutil.Reserse("Hello, World!"))
	structs := MyStruct.ExportStruct{Value: 7}
	fmt.Println(structs.PrintValue())
	fmt.Println(MyStruct.PrintVal())
}

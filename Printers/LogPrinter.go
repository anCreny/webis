package Printers

import "fmt"

func ShowLog(message ...any) {
	var prefix = "[" + colorBlue + "log" + colorReset + "]:"
	fmt.Print(prefix)
	fmt.Println(message...)
}

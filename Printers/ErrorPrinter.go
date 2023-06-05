package Printers

import "fmt"

func ShowError(message ...any) {
	var prefix = "[" + colorRed + "error" + colorReset + "]:"
	fmt.Print(prefix)
	fmt.Println(message...)
}

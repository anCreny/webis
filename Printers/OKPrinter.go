package Printers

import "fmt"

func ShowOk(message ...any) {
	var prefix = "[" + colorGreen + "OK" + colorReset + "]:"
	fmt.Print(prefix)
	fmt.Println(message...)
}

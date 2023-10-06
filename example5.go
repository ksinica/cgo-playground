package playground

// #include "example5.h"
import "C"
import (
	"fmt"
)

//export printFibGo
func printFibGo(n C.int) {
	a, b := 0, 1
	c := b
	for i := 0; ; i++ {
		c = b
		b = a + b
		if b >= int(n) {
			fmt.Println()
			break
		}
		a = c

		if i > 0 {
			fmt.Printf(",")
		}
		fmt.Printf("%d", b)
	}
}

func PrintFib(n int) {
	C.printFib(C.int(n))
}

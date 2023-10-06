package main

// #include <stdio.h>
//
// void printPtrAddr(void* ptr) {
//     printf("Address as seen by C:  %p\n", ptr);
// }
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

type foo struct {
	bar *int
}

func main() {
	var p runtime.Pinner

	v := 123

	C.printPtrAddr(unsafe.Pointer(&v))
	fmt.Printf("Address as seen by Go: %p\n", &v)

	f := foo{bar: &v}
	p.Pin(f.bar) // https://github.com/golang/go/issues/62380
	// Without pinning:
	// 	panic: runtime error: cgo argument has Go pointer to unpinned Go pointer
	C.printPtrAddr(unsafe.Pointer(&f))
	fmt.Printf("Address as seen by Go: %p\n", &f)

	p.Unpin()
}

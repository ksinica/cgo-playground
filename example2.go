package playground

// #include <stdio.h> // puts
// #include <stdlib.h> // free
import "C"
import "unsafe"

func Puts(str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C.puts(cstr)
}

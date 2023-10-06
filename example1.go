package playground

// int add(int a, int b) {
//     return a + b;
// }
//
import "C"

func Add(a, b int) int {
	return a + b
}

func AddCGO(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}

package playground

// #include <stdbool.h>
// #include <stdlib.h>
//
// #define SWAP(T, a, b) do { T c = a; a = b; b = c; } while (0)
//
// void bubbleSort(double arr[], int n) {
//     int i, j;
//     bool swapped;
//     for (i = 0; i < n - 1; i++) {
//         swapped = false;
//         for (j = 0; j < n - i - 1; j++) {
//             if (arr[j] > arr[j + 1]) {
//                  SWAP(double, arr[j], arr[j + 1]);
//                  swapped = true;
//             }
//         }
//
//         if (swapped == false) {
//             break;
//         }
//     }
// }
import "C"
import (
	"unsafe"
)

func BubbleSortFast(data []float64) {
	C.bubbleSort((*C.double)(unsafe.SliceData(data)), C.int(len(data)))
}

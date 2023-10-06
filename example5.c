#include "example5.h"
//#include "_cgo_export.h"

extern void printFibGo(int);

void printFib(int n) {
    printFibGo(n);
}
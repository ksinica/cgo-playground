package main

import (
	"fmt"
	"math/rand"

	playground "github.com/ksinica/cgo-playground"
)

func randFloat64s(n int) (ret []float64) {
	for n > 0 {
		ret = append(ret, rand.Float64())
		n--
	}
	return
}

func main() {
	data := randFloat64s(15)

	playground.BubbleSortFast(data)

	for i := range data {
		fmt.Println(data[i])
	}
}

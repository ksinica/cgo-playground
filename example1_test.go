package playground_test

import (
	"testing"

	playground "github.com/ksinica/cgo-playground"
)

func benchmarkExample1Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		playground.Add(i, i)
	}
}

func benchmarkExample1AddCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		playground.AddCGO(i, i)
	}
}

func BenchmarkExample1(b *testing.B) {
	b.Run("Add", benchmarkExample1Add)
	b.Run("AddCGO", benchmarkExample1AddCGO)
}

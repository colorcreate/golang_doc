package benchmark

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

func BenchmarkPerm(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Perm(5)
	}
}

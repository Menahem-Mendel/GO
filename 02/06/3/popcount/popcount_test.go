package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountTwo(0x1234567890ABCDEF)
	}
}
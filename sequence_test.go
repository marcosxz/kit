package kit

import (
	"testing"
)

func BenchmarkRandomCode(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		code := RandomCode(10)
		b.Log(string(code))
	}
}

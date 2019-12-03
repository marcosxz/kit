package kit

import (
	"testing"
)

var testCallFunction = func(arg1 int, arg2 bool) (int, bool) { return arg1, arg2 }

func BenchmarkCallFunction(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < 2; i++ {
		if results, err := CallFunction(testCallFunction, 1, true); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			for _, result := range results {
				b.Log(result)
			}
		}
	}
}

package kit

import (
	"bytes"
	"testing"
)

func BenchmarkCaller(b *testing.B) {
	b.ReportAllocs()
	var buffer = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		if _, file, _, ok := Caller(1); ok {
			buffer.WriteString(file + "\n")
		} else {
			b.Error("get caller error")
			b.FailNow()
		}
	}
	b.Log(buffer.String())
}

func BenchmarkCallerFrame(b *testing.B) {
	b.ReportAllocs()
	var buffer = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		if frame, ok := CallerFrame(1); ok {
			buffer.WriteString(frame.File + ":" + frame.Function + "\n")
		} else {
			b.Error("get caller frame error")
			b.FailNow()
		}
	}
	b.Log(buffer.String())
}

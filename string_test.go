package kit

import (
	"testing"
)

func BenchmarkStringToBytesAndBytesToString(b *testing.B) {

	b.ReportAllocs()

	s := "hdiasydkkashdkasjdkhkjashdasjdhasd"
	for i := 0; i < b.N; i++ {
		bs := StringToBytes(s)
		if len(bs) == 0 {
			b.Error("StringToBytes error")
			b.FailNow()
		}
		bss := BytesToString(bs)
		if bss != s {
			b.Error("BytesToString error")
			b.FailNow()
		}
	}
}

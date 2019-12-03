package kit

import (
	"testing"
)

func BenchmarkIdCardSex(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if sex, err := IdCardSex("510711197001071817"); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			b.Log(sex)
		}
	}
}

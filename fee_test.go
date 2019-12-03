package kit

import "testing"

func BenchmarkBranchToYuan(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.Log(BranchToYuan(14546))
	}
}

func BenchmarkYuanToBranch(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if branch, err := YuanToBranch("145.467"); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			b.Log(branch)
		}
	}
}

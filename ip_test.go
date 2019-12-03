package kit

import (
	"testing"
	"time"
)

func BenchmarkLocalIP(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if ip, err := LocalIP(); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			b.Log(ip)
		}
	}
}

func BenchmarkLocalIPByTcpDial(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if ip, err := LocalIPByTcpDial("www.baidu.com:80", time.Second*3); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			b.Log(ip)
		}
	}
}

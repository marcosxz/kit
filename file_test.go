package kit

import (
	"os"
	"testing"
)

var testFileDir = "testdata/"
var testFilePath = testFileDir + "test.file"

func BenchmarkFileIsExist(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		FileIsExist(testFilePath)
	}
}

func BenchmarkMakeDir(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := MakeDir(testFileDir); err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkCreateFile(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := CreateFile(testFilePath); err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkOpenFile(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if file, err := OpenFile(testFilePath, 0, 0777); err != nil {
			b.Error(err)
			b.FailNow()
		} else {
			b.Log(file.Name())
		}
	}
}

func BenchmarkReadDir(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := ReadDir(testFileDir, func(info os.FileInfo) bool {
			b.Log(info.Name())
			return true
		}); err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRemoveFile(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := RemoveFile(testFilePath); err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkRemoveDir(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if err := RemoveDir(testFileDir); err != nil {
			b.Error(err)
			b.FailNow()
		}
	}
}

func BenchmarkCurrentDir(b *testing.B) {

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.Log(CurrentDir())
	}
}

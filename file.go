package kit

import (
	"os"
	"path/filepath"
	"strings"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func MakeDir(path string) error {
	if !FileIsExist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func CreateFile(path string) error {
	path = strings.Replace(path, "\\", "/", -1)
	dir := path[:strings.LastIndex(path, "/")]
	if err := MakeDir(dir); err != nil {
		return err
	}
	file, err := os.Create(path)
	if err == nil {
		_ = file.Close()
	}
	return err
}

func OpenFile(path string, flag int, perm os.FileMode) (file *os.File, err error) {
	if !FileIsExist(path) {
		if err := CreateFile(path); err != nil {
			return nil, err
		}
	}
	return os.OpenFile(path, flag, perm)
}

func ReadDir(path string, reader func(os.FileInfo) bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	files, err := file.Readdir(-1)
	_ = file.Close()
	if err != nil {
		return err
	}
	for _, info := range files {
		if reader(info) {
			break
		}
	}
	return nil
}

func RemoveFile(path string) error {
	if FileIsExist(path) {
		return os.Remove(path)
	}
	return nil
}

func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

func CurrentDir() (path string, err error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1), err
}

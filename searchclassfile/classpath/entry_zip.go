package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry ZIP或JAR文件形式的类路径
type ZipEntry struct {
	absPath string //存放ZIP或JAR文件的绝对路径
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) //打开ZIP文件
	if err != nil {
		return nil, nil, err
	}
	defer r.Close() //defer语句来确保打开的文件得以关闭
	//遍历ZIP压缩包里的文件，挨个寻找指定的class文件
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open() //找到class文件，打开文件
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc) //读取class文件的所有数据
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self *ZipEntry) toString() string {
	return self.absPath
}

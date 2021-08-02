package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//DirEntry 目录形式的类路径
type DirEntry struct {
	absDir string //存放目录的绝对路径
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) //返回绝对路径
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) //拼接绝对路径和class文件名
	data, err := ioutil.ReadFile(fileName)            //读取class文件数据
	return data, self, err
}
func (self *DirEntry) toString() string {
	return self.absDir
}

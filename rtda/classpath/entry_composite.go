package classpath

import (
	"errors"
	"strings"
)

//CompositeEntry 多个目录或者文件形式的类路径
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	//根据分隔符分割为多个子路径，挨个创建对应的子路径的Entry实例
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	//依次调用每一个子路径的readClass方法
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) toString() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.toString()
	}
	return strings.Join(strs, pathListSeparator)
}

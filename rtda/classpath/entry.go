package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

// Entry /*定义一个接口Entry，用来表示类路径项*/
type Entry interface {
	/* readClass() 方法负责寻找和加载class文件*/
	readClass(className string) ([]byte, Entry, error)
	/* toString()方法用来输出字符串格式信息*/
	toString() string
}

// newEntry() 该方法根据不同的参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	/*指定的路径包含分隔符，也就是说类路径中同时包含多个文件或者目录*/
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	/*使用通配符符 * 指定某个目录下的所有文件*/
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	/*类路径指定的是jar文件或者zip文件*/
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	/*类路径是目录形式*/
	return newDirEntry(path)
}

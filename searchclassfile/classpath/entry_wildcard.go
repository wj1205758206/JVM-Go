package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] //移除末尾的通配符 *
	//WildcardEntry实际上也是CompositeEntry
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//通配符类路径不支持递归匹配子目录下的JAR文件，跳过子目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		//挑选出JAR文件
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	//返回的compositeEntry只包含通配符类路径下JAR文件，不包含子目录下的JAR文件
	return compositeEntry
}

package filedir

import (
	"os"
	"path"
)

// 是否存在
func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// 是否为文件
func isFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// 文件是否存在
func IsFileExist(f string) bool {
	if isExist(f) {
		return isFile(f)
	}
	return false
}

// 文件夹是否存在
func IsDirExist(f string) bool {
	if isExist(f) {
		return !isFile(f)
	}
	return false
}

// 获取文件名称(无路径,后缀)
func GetFileName(dst string) string {
	filenameall := path.Base(dst)
	filesuffix := path.Ext(dst)
	return filenameall[0 : len(filenameall)-len(filesuffix)]
}

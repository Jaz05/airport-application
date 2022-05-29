package path

import (
	"path/filepath"
	"runtime"
)

var basepath string
var appName = "airport-application"

func GetRootPath() string {
	if len(basepath) == 0 {
		basepath = findRootPath()
	}
	return basepath
}

func findRootPath() string {
	var _, path, _, _ = runtime.Caller(0)
	return findAppRoot(path)
}

func findAppRoot(path string) string {
	var dir = filepath.Dir(path)
	var dirName = filepath.Base(dir)
	if dirName != appName {
		return findAppRoot(dir)
	}
	return dir
}

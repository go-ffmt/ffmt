package ffmt

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 打印从当前行开始的栈
func MarkStackFull() {
	for i := 1; ; i++ {
		s := makeStackFunc(i)
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}

// 标记当前行栈
func MarkStack(skip int, a ...interface{}) {
	fmt.Print(makeStack(skip + 1))
	fmt.Println(a...)
}

// 标记当前行
func Mark(a ...interface{}) {
	MarkStack(1, a...)
}

var curDir = getCurrentDirectory()
var isTest = isTestd()

func isTestd() bool {
	return strings.Index(curDir, os.TempDir()) == 0
}

func getCurrentDirectory() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}

func getRelativeDirectory(basepath, targpath string) string {
	targpath = filepath.Clean(targpath)
	if isTest && strings.Index(targpath, runtime.GOROOT()) != 0 {
		return filepath.Base(targpath)
	}

	fileName, _ := filepath.Rel(basepath, targpath)
	if len(fileName) < len(targpath) {
		targpath = fileName
	}

	return targpath
}

func makeStack(skip int) string {
	_, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	fileName = getRelativeDirectory(curDir, fileName)
	return fmt.Sprintf("%s:%d ", fileName, line)
}

func makeStackFunc(skip int) string {
	pc, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Base(funcName)
	fileName = getRelativeDirectory(curDir, fileName)
	return fmt.Sprintf("%s:%d %s ", fileName, line, funcName)
}

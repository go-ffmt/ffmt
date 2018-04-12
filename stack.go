package ffmt

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// MarkStackFull Output stack full
func MarkStackFull() {
	for i := 1; ; i++ {
		s := FMakeStackFunc(i)
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}

// MarkStack Output prefix stack line pos
func MarkStack(skip int, a ...interface{}) {
	fmt.Println(append([]interface{}{FMakeStack(skip + 1)}, a...)...)
}

// Mark Output prefix current line position
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

// FMakeStack stack information
func FMakeStack(skip int) string {
	_, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	fileName = getRelativeDirectory(curDir, fileName)
	return fmt.Sprintf("%s:%d ", fileName, line)
}

// FMakeStackFunc stack information
func FMakeStackFunc(skip int) string {
	pc, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Base(funcName)
	fileName = getRelativeDirectory(curDir, fileName)
	return fmt.Sprintf("%s:%d %s ", fileName, line, funcName)
}

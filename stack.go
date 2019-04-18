package ffmt

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// MarkStackFull Output stack full
func MarkStackFull() {
	for i := 1; ; i++ {
		s := SmarkStackFunc(i)
		if s == "" {
			break
		}
		fmt.Println(s)
	}
}

// MarkStack Output prefix stack line pos
func MarkStack(skip int, a ...interface{}) {
	fmt.Println(append([]interface{}{SmarkStack(skip + 1)}, a...)...)
}

// Mark Output prefix current line position
func Mark(a ...interface{}) {
	MarkStack(1, a...)
}

// Smark returns Output prefix current line position
func Smark(a ...interface{}) string {
	return SmarkStack(1, a...)
}

var curDir, _ = os.Getwd()

func getRelativeDirectory(targpath string) string {
	targpath = filepath.Clean(targpath)

	if fileName, err := filepath.Rel(curDir, targpath); err == nil && len(fileName) <= len(targpath) {
		targpath = fileName
	}

	return targpath
}

// SmarkStack stack information
func SmarkStack(skip int, a ...interface{}) string {
	_, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	fileName = getRelativeDirectory(fileName)
	return fmt.Sprintf("%s:%d %s", fileName, line, fmt.Sprint(a...))
}

// SmarkStackFunc stack information
func SmarkStackFunc(skip int, a ...interface{}) string {
	pc, fileName, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return ""
	}
	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Base(funcName)
	fileName = getRelativeDirectory(fileName)
	return fmt.Sprintf("%s:%d %s %s", fileName, line, funcName, fmt.Sprint(a...))
}

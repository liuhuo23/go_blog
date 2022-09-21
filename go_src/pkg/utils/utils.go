package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// if 模拟 简单的三元操作

func If(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}

//GetRunPath 获取执行目录作为默认目录
func GetRunPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		return ""
	}
	return currentPath
}

// GetFileDirectoryToCaller 根据运行堆栈信息表获取文件目录， skip 默认1
func GetFileDirectoryToCaller(opts ...int) (directory string, ok bool) {
	var filename string
	directory = ""
	skip := 1

	if opts != nil {
		skip = opts[0]
	}
	if _, filename, _, ok = runtime.Caller(skip); ok {
		directory = path.Dir(filename)
	}
	return
}

// GetCurrentAbPathByExecutabel 获取当前执行文件的绝对路径
func GetCurrentAbPathByExecutable() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	res, _ := filepath.EvalSymlinks(exePath)
	return filepath.Dir(res), nil
}

//GetCurrentPath 获取当前执行文件路径，如果是临时目录则获取当前文件的执行路径
func GetCurrentPath() (dir string, err error) {
	dir, err = GetCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}
	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}
	if strings.Contains(dir, tmpDir) {
		var ok bool
		if dir, ok = GetFileDirectoryToCaller(2); !ok {
			return "", errors.New("failed to get path")
		}
	}
	return dir, nil
}

// GetDefaultPath 获取当前执行文件路径，如果临时目录则获取运行命令工作目录
func GetDefaultPath() (dir string, err error) {
	dir, err = GetCurrentAbPathByExecutable()
	if err != nil {
		return "", err
	}

	tmpDir, err := filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		return "", err
	}

	if strings.Contains(dir, tmpDir) {
		return GetRunPath(), nil
	}
	return dir, nil
}

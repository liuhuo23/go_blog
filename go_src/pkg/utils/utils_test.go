package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestGetRunPath(t *testing.T) {
	path := GetRunPath()
	print(path)
	if path == "" {
		t.Error("获取运行路径失败")
	}
}

func TestGetCurrentPath(t *testing.T) {
	path, err := GetCurrentPath()
	print(path)
	if err != nil {
		t.Error("获取运行路径失败")
	}
}

func TestGetCurrentAbPathByExecutable(t *testing.T) {
	_, err := GetCurrentAbPathByExecutable()
	if err != nil {
		t.Error("获取路径失败")
	}
}

func TestGetCurrentFileDirectory(t *testing.T) {
	path, ok := GetFileDirectoryToCaller()
	if !ok {
		t.Error("获取路径失败", path)
	}

	path, ok = GetFileDirectoryToCaller(1)
	if !ok {
		t.Error("获取路径失败", path)
	}
}

func TestIf(t *testing.T) {
	if 3 != If(false, 1, 3) {
		t.Error("模拟三元操作失败")
	}

	if 1 != If(true, 1, 3) {
		t.Error("模拟三元操作失败")
	}
}

func TestMyDecrypt(t *testing.T) {
	key := []byte("12345678")
	iv := []byte("12345678")
	data := []byte("12345678")
	out, err := MyEncrypt(data, key, iv)
	if err != nil {
		fmt.Println("err", err.Error())
		t.Error(err.Error())
	}

	fmt.Println(out)
	deByte := base64Encode(out)
	fmt.Println("deByte:", deByte)
	fmt.Println("hello")
	fmt.Println("加密后：", string(deByte))

	src := []byte("CyqS6B+0nOGkMmaqyup7gQ==")
	enByte, _ := base64Decode(src)
	out, _ = MyDecrypt(enByte, key, iv)
	log.Println("解密后:", string(out)) //2022/01/17 15:44:23 解密后: hello world
}

/*
* @Time    : 2021/3/15 20:36
* @Author  : CoderCharm
* @File    : read_file_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

读取本地文件操作

读取文件大小
https://stackoverflow.com/questions/17133590/how-to-get-file-length-in-go
**/
package _1_file_size

import (
	"os"
	"testing"
)

// 不打开文件直接使用 os.Stat 效率更高。
func readFileSize(path string) (int64, error) {
	file, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	// 获取文件大小
	size := file.Size()
	return size, nil
}

// 打开文件
func readFileSize2(path string) (int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// 再调用文件资源具柄对象
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	// 获取文件大小
	return fi.Size(), nil
}

func TestReadFile(t *testing.T) {
	fileSize, err := readFileSize("./foo.text")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(fileSize)
	}
}

func TestRead2File(t *testing.T) {
	fileSize, err := readFileSize2("./foo.text")
	if err != nil {
		t.Log(err)
	} else {
		t.Log(fileSize)
	}
}

/*
* @Time    : 2021/3/15 20:50
* @Author  : CoderCharm
* @File    : open_part_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

分段打开文件

https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go

**/
package _1_file_size

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func openFile(filePath string) {

	f, err := os.Open(filePath)
	if err != nil {
		panic(fmt.Sprintf("文件打开错误, %s", err))
	}
	// 最后关闭文件
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Sprintf("文件关闭错误, %s", err))
		}
	}()

	// 读取全部文件
	//content, err := ioutil.ReadAll(f)
	//fmt.Println(string(content))

	// 读取全部文件
	file, _ := os.Stat(filePath)
	// 获取文件大小
	allSize := file.Size()

	// 分段文件大小 5个字节
	partSize := 5
	// 计算全部分段
	partNum := math.Ceil(float64(allSize / int64(partSize)))

	// 分段读取文件
	for i := 0; i <= int(partNum); i++ {
		buf := make([]byte, partSize)

		byteLen, _ := f.Read(buf)

		fmt.Println(byteLen, string(buf))
	}

	//// 文件指针
	//_, _ = f.Seek(5, 1)
	//_, _ = f.Seek(5, 1)

}

func TestOpenFile(t *testing.T) {
	openFile("./foo.text")
}

/*
* @Time    : 2020-12-28 20:56
* @Author  : CoderCharm
* @File    : req4_down_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package http_demo

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func downloadFile(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}

type Reader struct {
	io.Reader
	Total   int64
	Current int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)

	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current*10000/r.Total)/100)

	return
}

func DownloadFileProgress(url, filename string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := &Reader{
		Reader: r.Body,
		Total:  r.ContentLength,
	}

	_, _ = io.Copy(f, reader)
}

func SoDownloadFile(url, filepath string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func TestDownFile(t *testing.T) {

	url := "https://bigota.d.miui.com/V12.0.5.0.QGHCNXM/miui_PHOENIX_V12.0.5.0.QGHCNXM_fc51e1211b_10.0.zip"
	filename := "miui.zip"

	_ = SoDownloadFile(url, filename)

}

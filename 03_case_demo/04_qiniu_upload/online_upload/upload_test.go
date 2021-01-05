/*
* @Time    : 2021-01-04 18:01
* @Author  : CoderCharm
* @File    : upload_test.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :

七牛云上传 在线视频
**/
package online_upload

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/parnurzeal/gorequest"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"strings"
	"testing"
)

const (
	Bucket     = ""
	AccessKey  = ""
	SecretKey  = ""
	DomainPath = ""
)

func QiniuUpLoad(targetUrl string) (string, error) {

	// 请求资源
	_, body, err := gorequest.New().Get(targetUrl).End()
	if err != nil {
		return "", errors.New(fmt.Sprintf("URL: %s 请求错误 原因: %s", targetUrl, err))
	}
	// 上传项目文件名
	suffix := getSuffixByUrl(targetUrl)

	//key := "aaa/bbb/go_qiniu_test2.mp4"
	key := fmt.Sprintf("%s.%s", GetUUID(), suffix)

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", Bucket, key),
	}

	// 认证
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	// 字节上传
	data := []byte(body)
	dataLen := int64(len(data))

	// 上传相关配置 是否启用https 选择上传机房 是否使用CDN上传加速 等配置 自行见文档
	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	// 上传
	qiniuErr := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if qiniuErr != nil {
		return "", errors.New(fmt.Sprintf("URL: %s 七牛云上传错误 原因: %s", targetUrl, err))
	}
	return fmt.Sprintf("%s/%s\n", DomainPath, ret.Key), nil
}

// 通过url截取 资源文件后缀
func getSuffixByUrl(targetUrl string) string {
	suffix := strings.Split(targetUrl, "?")
	suffix = strings.Split(suffix[0], ".")
	return suffix[len(suffix)-1]
}

// 获取uuid
func GetUUID() string {
	return uuid.New().String()
}

func TestUpload(t *testing.T) {

	mp4Url := "https://xxx.xxx.com/mv/0dd06bdcdf1e5e7f0d1988c21cfa91fe949dadf2.mp4"

	newUrl, err := QiniuUpLoad(mp4Url)

	fmt.Println(newUrl, err)

}

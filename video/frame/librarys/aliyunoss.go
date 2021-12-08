package librarys

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

type AliyunOss struct {
	Endpoint   string
	KeyID      string
	KeySecret  string
	BucketName string
}

const OssUrl = "http://tb-live.oss-cn-beijing.aliyuncs.com/"

func (u *AliyunOss) Upload(path, name string, data []byte) (string, bool) {
	client, err := oss.New(u.Endpoint, u.KeyID, u.KeySecret)
	if err != nil {
		logrus.Errorln("[srv.app]ossInfo err:", err)
		return "", false
	}
	bucket, err := client.Bucket(u.BucketName)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		return "", false
	}
	file := fmt.Sprintf("%s%s", path, name)
	// 上传Byte数组。
	err = bucket.PutObject(file, bytes.NewReader(data))
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		fmt.Println(err.Error())
		return "", false
	}
	return file, true
}

func (u *AliyunOss) UploadFile(path, name string, data io.Reader) (string, bool) {

	client, err := oss.New(u.Endpoint, u.KeyID, u.KeySecret)
	if err != nil {
		logrus.Errorln("[srv.app]ossInfo err:", err)
		return "", false
	}
	bucket, err := client.Bucket(u.BucketName)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		return "", false
	}
	if strings.LastIndex(path, "/") == -1 {
		path = path + "/"
	}
	file := fmt.Sprintf("%s%s", path, name)
	// 上传Byte数组。
	err = bucket.PutObject(file, data)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		fmt.Println(err.Error())
		return "", false
	}
	return file, true
}

func (o *AliyunOss) Delete(name string) bool {

	client, err := oss.New(o.Endpoint, o.KeyID, o.KeySecret)
	if err != nil {
		logrus.Errorln("[srv.app]ossInfo err:", err)
		return false
	}
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		return false
	}
	file := fmt.Sprintf("%s", name)
	err = bucket.DeleteObject(file)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		return false
	}
	return true
}

//Deletes 批量删除oss文件
func (u *AliyunOss) Deletes(names []string) bool {

	client, err := oss.New(u.Endpoint, u.KeyID, u.KeySecret)
	if err != nil {
		logrus.Errorln("[srv.app]ossInfo err:", err)
		return false
	}
	bucket, err := client.Bucket(u.BucketName)
	if err != nil {
		logrus.Errorln("[srv.app]Bucket err:", err)
		return false
	}

	for index := 0; index < len(names); index += 3 {
		objects := []string{}
		objects = append(objects, fmt.Sprintf("%s", names[index]))

		if index+1 < len(names) {
			objects = append(objects, fmt.Sprintf("%s", names[index+1]))
		}

		if index+2 < len(names) {
			objects = append(objects, fmt.Sprintf("%s", names[index+2]))
		}

		if _, err := bucket.DeleteObjects(objects); err != nil {
			logrus.Errorln("[srv.app]Bucket err:", err)
			return false
		}
	}
	return true
}

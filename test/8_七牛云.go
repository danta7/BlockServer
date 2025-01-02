package main

import (
	"BlogServer/core"
	"BlogServer/flags"
	"BlogServer/global"
	file2 "BlogServer/utlis/file"
	"BlogServer/utlis/hash"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/qiniu/go-sdk/v7/storagev2/credentials"
	"github.com/qiniu/go-sdk/v7/storagev2/http_client"
	"github.com/qiniu/go-sdk/v7/storagev2/uploader"
	"github.com/qiniu/go-sdk/v7/storagev2/uptoken"
	"io"
	"time"
)

func SendFile(file string) (url string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)

	hashString, err := hash.FileMd5(file)
	if err != nil {
		return
	}
	suffix, _ := file2.ImageSuffixJudge(file)
	fileName := fmt.Sprintf("%s.%s", hashString, suffix)
	key := fmt.Sprintf("%s/%s", global.Config.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadFile(context.Background(), file, &uploader.ObjectOptions{
		BucketName: global.Config.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   fileName,
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func SendReader(reader io.Reader) (url string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)

	uid := uuid.New().String()

	fileName := fmt.Sprintf("%s.png", uid)
	key := fmt.Sprintf("%s/%s", global.Config.QiNiu.Prefix, fileName)
	uploadManager := uploader.NewUploadManager(&uploader.UploadManagerOptions{
		Options: http_client.Options{
			Credentials: mac,
		},
	})
	err = uploadManager.UploadReader(context.Background(), reader, &uploader.ObjectOptions{
		BucketName: global.Config.QiNiu.Bucket,
		ObjectName: &key,
		FileName:   fileName,
	}, nil)
	return fmt.Sprintf("%s/%s", global.Config.QiNiu.Uri, key), err
}

func GenToken() (token string, err error) {
	mac := credentials.NewCredentials(global.Config.QiNiu.AccessKey, global.Config.QiNiu.SecretKey)

	putPolicy, err := uptoken.NewPutPolicy(global.Config.QiNiu.Bucket, time.Now().Add(1*time.Minute))
	if err != nil {
		return
	}
	token, err = uptoken.NewSigner(putPolicy, mac).GetUpToken(context.Background())
	if err != nil {
		return
	}
	return
}

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	//url, err := SendFile("uploads/images01/2331131388_董韬.jpg")
	//fmt.Println(url, err)
	fmt.Println(GenToken())
}

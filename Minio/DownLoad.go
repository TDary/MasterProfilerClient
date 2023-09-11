package Minio

import (
	"MasterClient/Logs"
	"context"

	"github.com/minio/minio-go/v7"
)

func DownLoadFile(objectName string, filePath string, contentType string) bool {
	ctx := context.Background()
	// 检查存储桶是否已经存在。
	exists, err := minioClient.BucketExists(ctx, RawBucketName)
	if err == nil && exists {
		Logs.Loggers().Printf("当前存储桶 %s存在----\n", RawBucketName)
	} else {
		Logs.Loggers().Printf("当前存储桶 %s不存在----\n", RawBucketName)
		Logs.Loggers().Print(err)
		return false
	}
	// 使用FGetObject下载文件。
	err = minioClient.FGetObject(ctx, RawBucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		Logs.Loggers().Println(err)
		return false
	}
	return true
}

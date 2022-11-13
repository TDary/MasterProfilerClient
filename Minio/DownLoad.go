package Minio

import (
	"MasterClient/Logs"

	"github.com/minio/minio-go/v6"
)

func DownLoadFile(objectName string, filePath string, contentType string, bucketName string) bool {
	// 检查存储桶是否已经存在。
	exists, err := minioClient.BucketExists(bucketName)
	if err == nil && exists {
		Logs.Loggers().Printf("We already own %s\n", bucketName)
	} else {
		Logs.Loggers().Fatalln(err)
	}
	Logs.Loggers().Printf("Successfully created %s\n", bucketName)
	// 使用FGetObject下载一个zip文件。
	err = minioClient.FGetObject(bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		Logs.Loggers().Println(err)
		return false
	}
	return true
}

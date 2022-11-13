package Minio

import (
	"MasterClient/Logs"

	"github.com/minio/minio-go/v6"
)

func UploadFile() {
	// 创建一个叫mymusic的存储桶。
	bucketName := "mymusic"
	location := "us-east-1"

	err := minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			Logs.Loggers().Printf("We already own %s\n", bucketName)
		} else {
			Logs.Loggers().Fatalln(err)
		}
	}
	Logs.Loggers().Printf("Successfully created %s\n", bucketName)

	// 上传一个zip文件。
	objectName := "golden-oldies.zip"
	filePath := "/tmp/golden-oldies.zip"
	contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		Logs.Loggers().Fatalln(err)
	}

	Logs.Loggers().Printf("Successfully uploaded %s of size %d\n", objectName, n)
}

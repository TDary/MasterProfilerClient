package Minio

import (
	"MasterClient/Logs"
	"context"

	"github.com/minio/minio-go/v7"
)

func UploadFile(objectName string, filePath string, contentType string) {
	//location := "us-east-1"
	ctx := context.Background()

	exists, err := minioClient.BucketExists(context.Background(), BucketName)
	if err == nil && exists {
		Logs.Loggers().Printf("存储桶%s存在----\n", BucketName)
	} else {
		Logs.Loggers().Printf("存储桶%s不存在----\n", BucketName)
		err := minioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{Region: "us-east-1", ObjectLocking: true}) //不存在 创建一个
		if err != nil {
			Logs.Loggers().Printf("存储桶%s创建失败----%s\n", BucketName, err.Error())
			return
		}
	}

	// 上传一个zip文件。
	// objectName := "golden-oldies.zip"
	// filePath := "/tmp/golden-oldies.zip"
	// contentType := "application/zip"

	// 使用FPutObject上传一个zip文件。
	n, err := minioClient.FPutObject(ctx, BucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		Logs.Loggers().Print(err)
		return
	}

	Logs.Loggers().Printf("上传%s至MInio成功,大小%d----\n", objectName, n.Size)
}

package Minio

import (
	"MasterClient/Logs"

	"github.com/minio/minio-go/v6"
)

func InitMinio() {
	endpoint := "play.min.io" //minio服务器url
	accessKeyID := "minio"
	secretAccessKey := "minio"
	useSSL := true

	// 初使化 minio client对象。
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		Logs.Loggers().Fatalln(err)
	}

	Logs.Loggers().Printf("%#v\n", minioClient) // minioClient初使化成功
}

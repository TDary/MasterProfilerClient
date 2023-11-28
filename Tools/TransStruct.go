package Tools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		data[field.Name] = v.Field(i).Interface()
	}
	return data
}

func MapToString(obj map[string]interface{}) string {
	dataType, _ := json.Marshal(obj)
	dataString := string(dataType)
	return dataString
}

// 使用AES对数据进行解密
func Decrypt(data, key []byte) ([]byte, error) { //密钥：eb3386a8a8f57a579c93fdfb33ec9471
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}

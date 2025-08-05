package Tools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
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
func Decrypt(ciphertext, key []byte) ([]byte, error) { //密钥：eb3386a8a8f57a579c93fdfb33ec9471
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 分离Nonce和密文
	nonce := ciphertext[:aes.BlockSize]
	actualCiphertext := ciphertext[aes.BlockSize:]

	// CTR流解密
	stream := cipher.NewCTR(block, nonce)
	plaintext := make([]byte, len(actualCiphertext))
	stream.XORKeyStream(plaintext, actualCiphertext)

	return plaintext, nil
}

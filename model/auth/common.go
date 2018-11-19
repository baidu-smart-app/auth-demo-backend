package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type Config struct {
	AppKey     string `json:"app_key"`
	SecrectKey string `json:"secrect_key"`
}

func Decrypt(ciphertext, sessionKey, iv, appKey string) (content string, err error) {
	key, err := base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return
	}

	ivbs, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	cipherbs, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return
	}

	var block cipher.Block

	if block, err = aes.NewCipher([]byte(key)); err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
	}

	cbc := cipher.NewCBCDecrypter(block, ivbs)
	cbc.CryptBlocks(cipherbs, cipherbs)

	cipherbs = pKCS7UnPadding(cipherbs, block.BlockSize())

	if len(cipherbs) < 20 {
		err = errors.New("bad content")
		return
	}

	// 前面16位可以直接抛弃，17-20表示明文长度
	// TODO
	size := cipherbs[19]

	if len(cipherbs) < 20+int(size) {
		err = errors.New("bad content")
		return
	}

	// 最后N位一定是appkey
	if string(cipherbs[size+20:]) != appKey {
		err = errors.New("illegal appkey")
		return
	}

	return string(cipherbs[20 : size+20]), nil
}

func pKCS7UnPadding(plantText []byte, blockSize int) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

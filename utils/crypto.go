package utils

import (
	"crypto/aes"
	"crypto/cipher"
)

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesCBCDecrypt(encryptData, key, iv []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	encryptData = encryptData[blockSize:]

	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)

	var decryptedData []byte
	mode.CryptBlocks(decryptedData, encryptData)
	decryptedData = PKCS7UnPadding(decryptedData)
	return decryptedData, nil
}

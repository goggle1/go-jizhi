package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	strMd5 := hex.EncodeToString(h.Sum(nil))
	return strMd5
}

func Encrypt(key, vector, src []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, fmt.Errorf("NewTripleDESCipher failure for key:%s, error:%v", key, err)
	}

	crypted, err := des3Encrypt(src, vector, block)
	if err != nil {
		return []byte{}, fmt.Errorf("des3AndBase64 des3Encrypt failure, error:%v", err)
	}
	res := base64Encode(crypted)

	return res, nil
}

func Decrypt(key, vector, src []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, fmt.Errorf("NewTripleDESCipher failure for key:%s, error:%v", key, err)
	}

	crypted, err := base64Decode(src)
	if err != nil {
		return []byte{}, fmt.Errorf("des3AndBase64 base64decode failure, error:%v", err)
	}

	res, err := des3Decrypt(crypted, vector, block)

	if err != nil {
		return []byte{}, fmt.Errorf("des3AndBase64 des3Decrypt failure, error:%v", err)
	}

	return res, nil
}

func base64Encode(src []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(buf, src)

	return buf
}

func base64Decode(src []byte) ([]byte, error) {
	enc := base64.StdEncoding

	dbuf := make([]byte, enc.DecodedLen(len(src)))
	n, err := enc.Decode(dbuf, src)
	return dbuf[:n], err
}

func des3Decrypt(crypted, vector []byte, block cipher.Block) (origData []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	blockMode := cipher.NewCBCDecrypter(block, vector)
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pkcs5UnPadding(origData)

	return
}

func des3Encrypt(originData, vector []byte, block cipher.Block) (crypted []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	origData := pkcs5Padding(originData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, vector)
	crypted = make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

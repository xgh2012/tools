/**
* @Author: XGH
* @Email: 55821284@qq.com
* @Date: 2021年12月28日15:49:09
*  AES ECB模式的加密解密
 */

package xencryp

import (
	"bytes"
	"crypto/aes"
	"errors"
)

// AesTool AES ECB模式的加密解密
type AesTool struct {
	//128 192  256位的其中一个 长度 对应分别是 16 24  32字节长度
	Key       []byte
	BlockSize int
}

func NewAesTool(key []byte, blockSize int) *AesTool {
	return &AesTool{Key: key, BlockSize: blockSize}
}

func (t *AesTool) padding(src []byte) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		//填充数据
		return append(src, bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)...)
	}
}

//unpadding
func (t *AesTool) unPadding(src []byte) []byte {
	length := len(src)
	padding := int(src[length-1])
	return src[:(length - padding)]
}

func (t *AesTool) Encrypt(src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(t.Key))
	if err != nil {
		return nil, err
	}
	//padding
	src = t.padding(src)
	//返回加密结果
	var encryptData []byte
	//存储每次加密的数据
	tmpData := make([]byte, t.BlockSize)

	//分组分块加密
	for index := 0; index < len(src); index += t.BlockSize {
		block.Encrypt(tmpData, src[index:index+t.BlockSize])
		encryptData = append(encryptData, tmpData...)
	}
	return encryptData, nil
}
func (t *AesTool) Decrypt(src []byte) ([]byte, error) {
	//避免崩溃
	srcLen :=len(src)
	if srcLen % t.BlockSize != 0{
		return nil,errors.New("数据格式不正确")
	}

	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(t.Key))
	if err != nil {
		return nil, err
	}
	//返回加密结果
	var decryptData []byte
	//存储每次加密的数据
	tmpData := make([]byte, t.BlockSize)

	//分组分块加密
	for index := 0; index < len(src); index += t.BlockSize {
		block.Decrypt(tmpData, src[index:index+t.BlockSize])
		decryptData = append(decryptData, tmpData...)
	}
	return t.unPadding(decryptData), nil
}
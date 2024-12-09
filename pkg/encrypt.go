/*
@Time : 2024/12/9 15:36
@Author : linx
@File : encrypt.go
@dsc: 加密
*/

package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

type Encrypt struct {
}

func NewEncrypt() *Encrypt {
	return &Encrypt{}
}

func (e *Encrypt) EncryptWithSalt(origin, salt string) string {
	combined := origin + salt
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

/*
@Time : 2024/12/9 15:31
@Author : linx
@File : random.go
@dsc: 随机数生成
*/

package pkg

import (
	"math/rand"
	"time"
)

type GenerateRandom struct {
}

func NewRandomGenerate() *GenerateRandom {
	return &GenerateRandom{}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateSalt 生成随机盐，这里假设生成一个指定长度的随机字符串作为盐，比如长度为16
func (g *GenerateRandom) GenerateSalt(saltLength int) string {
	saltLetters := []rune(letters)
	// 使用当前时间作为随机数生成器的种子，保证每次运行结果有变化
	// 使用当前时间作为随机数生成器的源，创建一个新的rand.Rand实例
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, saltLength)
	for i := range b {
		b[i] = saltLetters[r.Intn(len(saltLetters))]
	}
	return string(b)
}

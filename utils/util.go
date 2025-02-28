package utils

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnwoqrstwvuxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	// 创建一个新的随机数生成器
	seed := time.Now().UnixNano() // 使用当前时间作为种子
	rng := rand.New(rand.NewSource(seed))

	for i := range result {
		result[i] = letters[rng.Intn(len(letters))]
	}
	return string(result)
}

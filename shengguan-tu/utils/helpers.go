package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// GenerateID 生成唯一ID
func GenerateID(prefix string) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	timestamp := time.Now().UnixNano() / 1000000
	random := rnd.Intn(10000)
	return fmt.Sprintf("%s_%d_%04d", prefix, timestamp, random)
}

// Capitalize 首字母大写
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// Truncate 截断字符串
func Truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length-3] + "..."
}

// FormatNumber 格式化数字
func FormatNumber(n int) string {
	return fmt.Sprintf("%d", n)
}

// GetRandomString 获取随机字符串
func GetRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rnd.Intn(len(charset))]
	}
	return string(result)
}

// ContainsString 检查字符串是否在切片中
func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// RemoveString 从切片中移除字符串
func RemoveString(slice []string, s string) []string {
	for i, item := range slice {
		if item == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

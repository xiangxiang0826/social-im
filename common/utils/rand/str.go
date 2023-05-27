package randUtils

import (
	"math/rand"
	"strconv"
)

// RandPhone 随机手机号
func RandPhone() string {
	return "13" + RandNum(8)
}

// RandNum 随机数字
func RandNum(i int) string {
	var str string
	for j := 0; j < i; j++ {
		str += strconv.Itoa(RandInt(0, 9))
	}
	return str
}

// RandString returns a random string with a fixed length
func RandString(n int, allowedChars ...[]rune) string {
	var letters []rune
    if len(allowedChars) == 0 {
		letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	} else {
		letters = allowedChars[0]
	}
    b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
    return string(b)
}

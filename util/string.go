package util

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")

// RandomString get random string
func RandomString(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// GenUlid generater ulid
func GenUlid() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// Sha256 gen sha256 string
func Sha256(str string) string {
	hash := sha256.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashBytes)
}

// ContainsString check the string slice contain the string
func ContainsString(arr []string, str string) (found bool) {
	for _, v := range arr {
		if !found && v == str {
			found = true
		}
	}
	return
}

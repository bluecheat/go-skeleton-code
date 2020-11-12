package utils

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 랜덤 난수 생성기
func RandToken(n ...int) string {
	if n[0] == 0 {
		n[0] = 32
	}
	b := make([]rune, n[0])
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Hash(password string) (string, error) {
	hexPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hexPassword), err

}

func VerifyPassword(hash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			// MEMO: err를 wrap 하여 상세를 전달하면 좋다
			return err
		}
		return err
	}
	return nil
}

package datacookie

import (
	"math/rand"

	"github.com/go-chi/jwtauth"
)

var (
	tokenizer *jwtauth.JWTAuth
)

func init() {
	randomSecret := randomSecret()
	initializeTokenValue(randomSecret)
}

func initializeTokenValue(secret string) {
	tokenizer = jwtauth.New("HS256", []byte(secret), nil)
}

// func make random string with 32 symbols length
func randomSecret() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// SetDataCookieSecret set secret for data cookie
func SetDataCookieSecret(secret string) {
	initializeTokenValue(secret)
}

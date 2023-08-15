package sanctum

import (
	"fmt"
	"math/rand"
	"time"
)

type Hawaldar struct {
}

const length = 20

// This package will take care of generating personal access token.
// When the request comes through, it will verify the authorization
// of request and resolve bearer token to an identity (user).

func Resolve(token string) {

}

func Generate() string {
	return randomString()
}

func randomString() string {
	rand.NewSource(time.Now().UnixNano())
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	var random [length]rune
	for i := 0; i < length; i++ {
		random[i] = chars[rand.Int()%length]
	}
	return fmt.Sprintf("%c", random)
}

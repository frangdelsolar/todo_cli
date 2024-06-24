package helpers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/frangdelsolar/todo_cli/pkg/auth"
)

const (
    letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
    numBytes    = "0123456789"
)

func CreateRandomUser() (*auth.User, error) {
    name := RandomName()
    email := RandomEmail()
    password := RandomPassword()
    return auth.CreateUser(name, email, password)
}

func RandomName() string {
    return RandomString(10)
}

func RandomEmail() string {
    name := RandomString(8)
    domain := "example.com"
    return fmt.Sprintf("%s@%s", name, domain)
}

func RandomPassword() string {
    // Adjust these parameters as needed for your password complexity requirements
    minLength := 12
    maxLength := 16

    length := rand.Intn(maxLength - minLength) + minLength

    charPool := []rune{}
    charPool = append(charPool, []rune(letterBytes)...)
    charPool = append(charPool, []rune(numBytes)...)

    password := ""
    for i := 0; i < length; i++ {
        password += string(charPool[rand.Intn(len(charPool))])
    }

    return password
}

func RandomString(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := []byte("INI ADALAH TOKEN RAHASIA")
	claims := jwt.MapClaims{
		"uid":  "1234",
		"role": "admin",
		"exp":  time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	singned, err := token.SignedString(secret)
	if err != nil {
		panic(err)
	}
	fmt.Println("generate token:")
	fmt.Println(singned)
}

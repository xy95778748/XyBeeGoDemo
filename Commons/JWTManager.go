package Commons

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
)

const key = "hello world"

func GenerateToken() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		fmt.Println(err)
	}

	return tokenString, err
}
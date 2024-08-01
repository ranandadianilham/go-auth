package auth

import (
	"fmt"
	"time"

	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/golang-jwt/jwt/v5"
)

//var secretKey = []byte("secret")

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	privatekey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return privatekey, nil
}

func CreateToken(username string) (string, error) {
	privateKey, err := GenerateECDSAKey()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	fmt.Println("token", tokenString)
	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secret, _ := GenerateECDSAKey()
		return secret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

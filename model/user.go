package model

import (
	"github.com/dgrijalva/jwt-go"
	"os"

	// "fmt"
)

const (
	userIDKey = "user_id"
)

var (
	secret = os.Getenv("ADMIN_JWT_SECRET")
)

func VerifyAccessToken(tokenString string) (UserID, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", nil
		}
		return []byte(secret), nil
	})
	// if err != nil {
	// 	// ErrInvalidToken
	// 	fmt.Printf("%v", "err")
	// 	return "", err
	// }

	if token == nil {
		return "", nil
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return "", nil
	// }

	userID, _ := claims[userIDKey].(string)
	// if !ok {
	// 	return "", nil
	// }

	// verifiedPhoneNumber, ok := claims[verifiedPhoneNumberKey].(bool)
	// if !ok {
	// 	return "", false, nil
	// }

	return UserID(userID), nil
}

package utility

import (
	jwt "github.com/golang-jwt/jwt/v4"
	t "time"
	"fmt"
)

var secret = []byte("secret")

func GenerateToken(data map[string]interface{}) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"exp": jwt.NewNumericDate(t.Now().Add(t.Minute * 2)),
		"iat": jwt.NewNumericDate(t.Now()),
		"sub": data,
	})
	token, _ := claims.SignedString([]byte(secret))
	return token
}

func ParseToken(ts string) (map[string]interface{}, error) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method:")
		}
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["sub"].(map[string]interface{}), nil
	} else {
		return nil, err
	}
}

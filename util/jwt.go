package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//tokeno luşturma

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})

	return claims.SignedString([]byte(SecretKey))
}

func ParseJwt(cookie string) (string, error) {
	//cookie yi parse ediyoruz

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Issuer, err
}

func ParseJwtId(cookie string) (string, error) {
	//cookie yi parse ediyoruz

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(*jwt.StandardClaims)

	return claims.Id, err
}

/*func GetUserIDFromJWT(cookie string) (uint, error) {
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid JWT claims")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("Invalid user ID in JWT claims")
	}

	userID := uint(sub)

	return userID, nil
}
*/

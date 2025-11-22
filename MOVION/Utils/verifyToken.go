package utils

import (
	config "movion/Config"
	models "movion/Models"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenStr string)(*models.Claims, error){
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		return config.Jwtkey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
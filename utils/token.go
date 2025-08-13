package utils

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("chave-secreta")

func GerarToken(id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatUint(uint64(id), 10),
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtKey)

}

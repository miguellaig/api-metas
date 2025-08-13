package middleware

import (
	"fmt"
	"net/http"
	"projeto-metas/dto/common"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var jwtKey = []byte("chave-secreta")

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Header token ausente"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "Header token em formato inválido"})
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("algoritmo inválido: %v", token.Method.Alg())
			}
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "Token inválido"})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao ler o map claims"})
		}

		idStr, ok := claims["sub"].(string)
		if !ok {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "sub claims não é uma string"})
		}

		userIDUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao converter string para uint"})
		}

		userID := uint(userIDUint)

		c.Set("user", userID)

		return next(c)

	}
}

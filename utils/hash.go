package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GerarHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("falha ao gerar hash: %w", err)
	}

	return string(hash), nil
}

func CompararHashESenha(hash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}

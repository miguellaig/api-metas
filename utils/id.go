package utils

import (
	"errors"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetIDParamUint(c echo.Context, paramName string) (uint, error) {
	idParam := c.Param(paramName)
	if idParam == "" {
		return 0, errors.New("parâmetro ausente")
	}

	idUint64, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(idUint64), nil
}

func PegarUserID(c echo.Context) (uint, error) {
	userID, ok := c.Get("user").(uint)
	if !ok {
		return 0, errors.New("valor do user no contexto não é um uint")
	}

	return userID, nil

}

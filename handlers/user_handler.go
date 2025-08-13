package handlers

import (
	"errors"
	"net/http"
	"projeto-metas/dto/common"
	"projeto-metas/dto/request"
	"projeto-metas/services"
	"projeto-metas/transformer"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (u *UserHandler) RegistroUserHandler(c echo.Context) error {
	var input request.CreateUserInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: err.Error()})
	}

	user, err := u.service.RegistrarUsuario(input)
	if errors.Is(err, services.ErrEmailJaCadastrado) {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "email já cadastrado"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "erro interno"})
	}

	return c.JSON(http.StatusCreated, transformer.ToUserResponse(user))

}

func (u *UserHandler) LoginUserHandler(c echo.Context) error {
	var input request.LoginUser

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Falha ao ler o json (login user)"})
	}

	token, err := u.service.LoginUsuario(input.Email, input.Password)
	if err != nil {
		switch err {
		case services.ErrEmailNaoCadastrado:
			return c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "email não cadastrado"})
		case services.ErrSenhaInvalida:
			return c.JSON(http.StatusUnauthorized, common.ErrorResponse{Error: "senha inválida"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "erro interno"})
		}

	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})

}

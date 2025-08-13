package services

import (
	"errors"
	"projeto-metas/dto/request"
	"projeto-metas/models"
	"projeto-metas/repository"
	"projeto-metas/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

var ErrEmailJaCadastrado = errors.New("email já cadastrado")
var ErrEmailNaoCadastrado = errors.New("email não cadastrado")
var ErrSenhaInvalida = errors.New("senha inválida")

func (u *UserService) RegistrarUsuario(input request.CreateUserInput) (*models.User, error) {

	existente, err := u.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if existente != nil {
		return nil, ErrEmailJaCadastrado
	}

	hash, err := utils.GerarHash(input.Password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Email:    input.Email,
		Password: hash,
	}

	return u.repo.CreateUser(newUser)
}

func (u *UserService) LoginUsuario(email, senha string) (string, error) {
	existente, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if existente == nil {
		return "", ErrEmailNaoCadastrado
	}

	if err := utils.CompararHashESenha(existente.Password, senha); err != nil {
		return "", ErrSenhaInvalida
	}

	token, err := utils.GerarToken(existente.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

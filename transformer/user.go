package transformer

import (
	"projeto-metas/dto/request"
	"projeto-metas/dto/response"
	"projeto-metas/models"
)

func ToUserModel(input *request.CreateUserInput) *models.User {
	return &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

}
func ToUserResponse(input *models.User) *response.UserResponse {
	return &response.UserResponse{
		Name:   input.Name,
		Email:  input.Email,
		UserID: input.ID,
	}
}

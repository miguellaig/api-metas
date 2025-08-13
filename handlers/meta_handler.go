package handlers

import (
	"net/http"
	"projeto-metas/dto/common"
	"projeto-metas/dto/request"
	"projeto-metas/services"
	"projeto-metas/transformer"
	"projeto-metas/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type MetaHandler struct {
	service  *services.MetaService
	validate *validator.Validate
}

func NewMetaHandler(s *services.MetaService, v *validator.Validate) *MetaHandler {
	return &MetaHandler{
		service:  s,
		validate: v,
	}
}

func (m *MetaHandler) RegistrarMeta(c echo.Context) error {
	userID, err := utils.PegarUserID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "User salvo no contexto não é um uint"})
	}

	var meta request.MetaRequest

	if err := c.Bind(&meta); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Falha ao ler o json"})
	}

	if err := m.validate.Struct(meta); err != nil {
		msg := utils.FormatValidationError(err)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"errors": msg})
	}

	novaMeta, err := m.service.RegistrarMeta(userID, &meta)
	if err != nil {
		switch err {
		case services.ErrMetaJaExiste:
			return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Meta com mesmo título já existe"})
		case services.ErrInterno:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro interno"})
		case services.ErrFalhaAoCriarMeta:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao criar meta"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro inesperado"})
		}
	}

	metaResponse := transformer.ToMetaResponse(novaMeta)

	return c.JSON(http.StatusCreated, common.ApiResponse{Data: metaResponse})
}

func (m *MetaHandler) ListarMetas(c echo.Context) error {
	userID, err := utils.PegarUserID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "User salvo no contexto não é uint"})
	}

	metaList, err := m.service.ListarMetas(userID)
	if err != nil {
		switch err {
		case services.ErrFalhaAoListarMetas:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao listar metas"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro inesperado"})
		}
	}
	return c.JSON(http.StatusOK, common.ApiResponse{Data: metaList})
}

func (m *MetaHandler) ListarMetaPorID(c echo.Context) error {
	userID, err := utils.PegarUserID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "User salvo no contexto não é uint"})
	}
	id, err := utils.GetIDParamUint(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID da url incorreto"})
	}

	meta, err := m.service.ListarMetaPorID(userID, id)
	if err != nil {
		switch err {
		case services.ErrFalhaAoListarMetas:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao listar meta"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro inesperado"})
		}
	}
	return c.JSON(http.StatusOK, common.ApiResponse{Data: meta})
}

func (m *MetaHandler) UpdateMeta(c echo.Context) error {
	userID, err := utils.PegarUserID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "User salvo no contexto não é uint"})
	}
	id, err := utils.GetIDParamUint(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID da url incorreto"})
	}

	var meta request.MetaRequest

	if err := c.Bind(&meta); err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "Falha ao ler o json"})
	}

	newMeta, err := m.service.UpdateMetaPorID(userID, id, &meta)
	if err != nil {
		switch err {
		case services.ErrMetaNaoExiste:
			return c.JSON(http.StatusNotFound, common.ErrorResponse{Error: "Meta não existe"})
		case services.ErrFalhaAoAtualizarMeta:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao atualizar meta"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro inesperado"})
		}
	}
	return c.JSON(http.StatusOK, common.ApiResponse{Data: newMeta})
}

func (m *MetaHandler) DeleteMeta(c echo.Context) error {
	userID, err := utils.PegarUserID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "User salvo no contexto não é uint"})
	}
	id, err := utils.GetIDParamUint(c, "id")
	if err != nil {
		return c.JSON(http.StatusBadRequest, common.ErrorResponse{Error: "ID da url incorreto"})
	}

	err = m.service.DeleteMeta(userID, id)
	if err != nil {
		switch err {
		case services.ErrMetaNaoExiste:
			return c.JSON(http.StatusNotFound, common.ErrorResponse{Error: "Meta não existe"})
		case services.ErrFalhaAoDeletarUsuario:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Falha ao deletar meta"})
		default:
			return c.JSON(http.StatusInternalServerError, common.ErrorResponse{Error: "Erro inesperado"})
		}
	}
	return c.NoContent(http.StatusNoContent)
}

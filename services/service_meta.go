package services

import (
	"errors"
	"projeto-metas/dto/request"
	"projeto-metas/models"
	"projeto-metas/repository"
	"projeto-metas/transformer"
)

type MetaService struct {
	repo *repository.MetaRepository
}

func NewMetaService(repo *repository.MetaRepository) *MetaService {
	return &MetaService{repo}
}

var ErrInterno = errors.New("erro interno")
var ErrMetaJaExiste = errors.New("meta com mesmo título já existe")
var ErrFalhaAoCriarMeta = errors.New("falha ao criar nova meta")
var ErrFalhaAoListarMetas = errors.New("falha ao listar metas")
var ErrMetaNaoExiste = errors.New("meta não existe")
var ErrFalhaAoAtualizarMeta = errors.New("falha ao atualizar meta")
var ErrFalhaAoDeletarUsuario = errors.New("falha ao deletar usuário")

func (m *MetaService) RegistrarMeta(userid uint, metaReq *request.MetaRequest) (*models.Meta, error) {
	existingMeta, err := m.repo.FindByTitleAndUserID(metaReq.Titulo, userid)
	if err != nil {
		return nil, ErrInterno
	}
	if existingMeta != nil {
		return nil, ErrMetaJaExiste
	}

	meta := transformer.ToMetaModel(metaReq)
	meta.UserID = userid

	newMeta, err := m.repo.CreateMeta(meta)
	if err != nil {
		return nil, ErrFalhaAoCriarMeta
	}
	return newMeta, nil
}

func (m *MetaService) ListarMetas(id uint) ([]models.Meta, error) {
	metaList, err := m.repo.FindByUserID(id)
	if err != nil {
		return nil, ErrFalhaAoListarMetas
	}
	return metaList, nil
}

func (m *MetaService) ListarMetaPorID(userID, id uint) (*models.Meta, error) {
	meta, err := m.repo.FindBYIDAndUserID(id, userID)
	if err != nil {
		return nil, ErrFalhaAoListarMetas
	}
	return meta, nil
}

func (m *MetaService) UpdateMetaPorID(userID, id uint, newMeta *request.MetaRequest) (*models.Meta, error) {
	meta, err := m.repo.FindBYIDAndUserID(id, userID)
	if err != nil {
		return nil, ErrMetaNaoExiste
	}

	if err := m.repo.UpdateMeta(userID, id, newMeta); err != nil {
		return nil, ErrFalhaAoAtualizarMeta
	}

	meta.Titulo = newMeta.Titulo
	meta.Descricao = newMeta.Descricao
	meta.Status = newMeta.Status
	meta.Prazo = newMeta.Prazo

	return meta, nil
}
func (m *MetaService) DeleteMeta(userid, id uint) error {
	_, err := m.repo.FindBYIDAndUserID(id, userid)
	if err != nil {
		return ErrMetaNaoExiste
	}

	if err := m.repo.DeleteMeta(userid, id); err != nil {
		return ErrFalhaAoDeletarUsuario
	}

	return nil

}

package repository

import (
	"errors"
	"projeto-metas/dto/request"
	"projeto-metas/models"

	"gorm.io/gorm"
)

type MetaRepository struct {
	db *gorm.DB
}

func NewMetaRepository(db *gorm.DB) *MetaRepository {
	return &MetaRepository{db}
}

func (m *MetaRepository) FindByTitleAndUserID(title string, userid uint) (*models.Meta, error) {

	var meta models.Meta

	if err := m.db.Where("titulo = ? AND user_id = ?", title, userid).First(&meta).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &meta, nil

}
func (m *MetaRepository) FindByUserID(id uint) ([]models.Meta, error) {
	var meta []models.Meta

	if err := m.db.Where("user_id = ?", id).Find(&meta).Error; err != nil {
		return nil, err
	}
	return meta, nil
}
func (m *MetaRepository) FindBYIDAndUserID(id, userID uint) (*models.Meta, error) {
	var meta models.Meta

	if err := m.db.Where("user_id = ? AND id = ?", userID, id).First(&meta).Error; err != nil {
		return nil, err
	}
	return &meta, nil
}

func (m *MetaRepository) CreateMeta(meta *models.Meta) (*models.Meta, error) {
	if err := m.db.Create(&meta).Error; err != nil {
		return nil, err
	}
	return meta, nil
}

func (m *MetaRepository) UpdateMeta(userID, id uint, meta *request.MetaRequest) error {
	return m.db.Model(&models.Meta{}).Where("user_id = ? AND id = ?", userID, id).Updates(meta).Error
}
func (m *MetaRepository) DeleteMeta(userid, id uint) error {
	return m.db.Where("user_id = ? AND id = ?", userid, id).Delete(&models.Meta{}).Error
}

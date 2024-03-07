package repositories

import (
	"errors"
	"final_project_promotion/internal/app/models"
	"final_project_promotion/utils/exceptions"

	"gorm.io/gorm"
)

type PromotionRepository interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionRepositoryImpl struct {
	db *gorm.DB
}

// NewPromotionRepository creates a new instance of PromotionRepository
func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &PromotionRepositoryImpl{
		db: db,
	}
}

// CreatePromotion creates a new promotion in the database
func (r *PromotionRepositoryImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	err := r.db.Unscoped().Create(&promo).Error
	return promo, err
}

// GetAllPromotions throw all data that recorded in the database
func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	// Implementasi kamu taruh disini
	var promotions []models.Promotion
	if err := r.db.Debug().Unscoped().Find(&promotions).Error; err != nil {
		return nil, err
	}
	return promotions, nil
}

// GetPromotionByPromotionID will throw data based on promotionID request
// func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
// 	// Implementasi kamu taruh disini

// }

// GetPromotionByPromotionID will throw data based on promotionID request
func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
	var promo models.Promotion
	if err := r.db.Unscoped().Where("promotion_id = ?", PromotionID).Take(&promo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

			// Handle case where record is not found
			// For example, you can return a specific error indicating that the record is not found
			return models.Promotion{}, errors.New("promotion not found")
		}
		// Handle other errors
		return models.Promotion{}, err
	}
	return promo, nil
}

// UpdatePromotion will update data based on promotionID request
func (r *PromotionRepositoryImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Implementasi kamu taruh disini
	var exits models.Promotion
	if err := r.db.Where("promotion_id = ?", promo.PromotionID).First(&exits).Error; err != nil {
		return models.Promotion{}, err
	}
	// Update the promotion
	if err := r.db.Unscoped().Save(&promo).Error; err != nil {
		return models.Promotion{}, err
	}
	return promo, nil
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	// Implementasi kamu taruh disini
	if err := r.db.Unscoped().Where("promotion_id = ?", promotionID).Delete(&models.Promotion{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &exceptions.PromotionIDNotFoundError{
				Message: "Promotion Not Found",
				PromotionID: promotionID,
			}
		}
		return err
	}
	return nil
}
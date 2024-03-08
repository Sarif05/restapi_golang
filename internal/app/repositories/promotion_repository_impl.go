package repositories

import (
	"errors"
	models "final_project_promotion/internal/app/models/domain"
	"final_project_promotion/utils/exceptions"

	"gorm.io/gorm"
)

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
	err := r.db.Unscoped().Create(&promo).Error
	return promo, err
}

func (r *PromotionRepositoryImpl) GetAllPromotions(limit, offset int) ([]models.Promotion, error) {
	var promotions []models.Promotion
	if err := r.db.Debug().Unscoped().Limit(limit).Offset(offset).Find(&promotions).Error; err != nil {
		return nil, err
	}
	return promotions, nil
}

// GetPromotionByPromotionID will throw data based on promotionID request
func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
	var promotions models.Promotion
	if err := r.db.Unscoped().Where("promotion_id = ?", PromotionID).Take(&promotions).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Promotion{}, errors.New("promotion not found")
		}
		return models.Promotion{}, err
	}
	return promotions, nil
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

func (r *PromotionRepositoryImpl) SearchPromotions(query string, limit, offset int) ([]models.Promotion, error) {
    var promotions []models.Promotion
    if err := r.db.Debug().Unscoped().Where("(promotion_name LIKE ? OR discount_type LIKE ? OR CAST(discount_value AS TEXT) LIKE ? OR TO_CHAR(promotion_start_date, 'YYYY-MM-DD') LIKE ? OR TO_CHAR(promotion_end_date, 'YYYY-MM-DD') LIKE ?)", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").
        Limit(limit).Offset(offset).Find(&promotions).Error; err != nil {
        return nil, err
    }
    return promotions, nil
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	// Implementasi kamu taruh disini
	if err := r.db.Unscoped().Where("promotion_id = ?", promotionID).Delete(&models.Promotion{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound{
			return &exceptions.NotFoundErr{
				Message: "Promotion Not Found",
			}
		}
		return err
	}
	return nil
}

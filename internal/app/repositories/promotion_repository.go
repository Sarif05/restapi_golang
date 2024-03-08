package repositories

import (
	models "final_project_promotion/internal/app/models/domain"
)

// type PromotionRepository interface {
// 	CreatePromotion(promo models.Promotion) (models.Promotion, error)
// 	GetAllPromotions(limit, offset int) ([]models.Promotion, error)
// 	SearchPromotions(query string, limit, offset int) ([]models.Promotion, error)
// 	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
// 	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
// 	DeletePromotionbyPromotionID(promotionID string) error
// }

type PromotionRepository interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions(limit, offset int) ([]models.Promotion, error)
	SearchPromotions(query string, limit, offset int) ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}
package services

import (
	"final_project_promotion/internal/app/models"
)

type PromotionService interface{
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions(limit, offset int) ([]models.Promotion, error)
	SearchPromotions(query string, limit, offset int) ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}


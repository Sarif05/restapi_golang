package services

import (
	"final_project_promotion/internal/app/models/web"
)

type PromotionService interface{
	CreatePromotion(request web.PromotionCreateRequest) (web.PromotionResponse, error)
	GetAllPromotions(limit, offset int) ([]web.PromotionResponse, error)
	SearchPromotions(query string, limit, offset int) ([]web.PromotionResponse, error)
	GetPromotionbyPromotionID(promotionID string) (web.PromotionResponse, error)
	UpdatePromotionbyPromotionID(request web.PromotionUpdateRequest) (web.PromotionResponse, error)
	DeletePromotionbyPromotionID(promotionID string) error
}
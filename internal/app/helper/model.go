package helper

import (
	models "final_project_promotion/internal/app/models/domain"
	"final_project_promotion/internal/app/models/web"
)
func ToPromotionResponse(promotion models.Promotion) web.PromotionResponse {
	return web.PromotionResponse{
		Id:   int(promotion.ID),
		PromotionID: promotion.PromotionID,
		PromotionName: promotion.PromotionName,
		DiscountType: promotion.DiscountType,
		DiscountValue: promotion.DiscountValue,
		PromotionStartDate: promotion.PromotionStartDate,
		PromotionEndDate: promotion.PromotionEndDate,
		CreatedAt: promotion.CreatedAt,
		UpdatedAt: promotion.UpdatedAt,
	}
}

func ToPromotionResponses(promotions []models.Promotion) []web.PromotionResponse {
	var promotionResponses []web.PromotionResponse
	for _, promotion := range promotions {
		promotionResponses = append(promotionResponses, ToPromotionResponse(promotion))
	}
	return promotionResponses
}

package web

import "time"

type PromotionResponse struct {
	Id   int    `json:"id"`
	PromotionID string `json:"promotion_id"`
	PromotionName string `json:"promotion_name"`
	DiscountType string `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	PromotionStartDate time.Time `json:"promotion_start_date"`
	PromotionEndDate time.Time `json:"promotion_end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
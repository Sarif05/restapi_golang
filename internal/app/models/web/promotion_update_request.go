package web

import "time"

type PromotionUpdateRequest struct{
	Id int `validate="required"` 
	PromotionID string `validate:"required,min=1,max=100" json:"promotion_id"`
	PromotionName string `validate:"required,min=1,max=100" json:"promotion_name"`
	DiscountType string `validate:"required,min=1,max=100" json:"discount_type"`
	DiscountValue float64 `validate:"required" json:"discount_value"`
	PromotionStartDate time.Time `validate:"required" json:"promotion_start_date"`
	PromotionEndDate time.Time `validate:"required" json:"promotion_end_date"`
}
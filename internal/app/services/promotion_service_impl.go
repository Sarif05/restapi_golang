package services

import (
	"final_project_promotion/internal/app/helper"
	models "final_project_promotion/internal/app/models/domain"
	"final_project_promotion/internal/app/models/web"
	"final_project_promotion/internal/app/repositories"
	"final_project_promotion/utils/exceptions"
)

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

func NewPromotionService (PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl{
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

func (s *PromotionServiceImpl) CreatePromotion(promo web.PromotionCreateRequest) (web.PromotionResponse, error) {
	promotion := models.Promotion{
		PromotionName:       promo.PromotionName,
		DiscountType:       promo.DiscountType,
		DiscountValue:      promo.DiscountValue,
		PromotionStartDate: promo.PromotionStartDate,
		PromotionEndDate:   promo.PromotionEndDate,
	}

	createdPromotion, err := s.PromotionRepo.CreatePromotion(promotion)
	if err != nil {
		return web.PromotionResponse{}, err
	}

	return helper.ToPromotionResponse(createdPromotion), nil
}

func (s *PromotionServiceImpl) SearchPromotions(query string, limit, offset int) ([]web.PromotionResponse, error) {
    Promotion, err := s.PromotionRepo.SearchPromotions(query, limit, offset)
	if err != nil {
		return []web.PromotionResponse{}, err
	}
	return helper.ToPromotionResponses(Promotion), nil
}

func (s *PromotionServiceImpl) GetAllPromotions(limit, offset int) ([]web.PromotionResponse, error) {
	Promotion, err := s.PromotionRepo.GetAllPromotions(limit, offset)
	if err != nil {
		return []web.PromotionResponse{}, err
	}
	return helper.ToPromotionResponses(Promotion), nil
}

func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (web.PromotionResponse, error) {
	promotion, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		// Mengembalikan error langsung tanpa menciptakan instance dari models.Promotion yang tidak diperlukan.
		return web.PromotionResponse{}, &exceptions.PromotionIDNotFoundError{
			Message:     "Promotion Not Found",
			PromotionID: promotionID,
		}
	}
	response := helper.ToPromotionResponse(promotion)
	return response, nil
}

func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo web.PromotionUpdateRequest) (web.PromotionResponse, error) {
	existingPromotion, err := s.PromotionRepo.GetPromotionbyPromotionID(promo.PromotionID)
	if err != nil {
		return web.PromotionResponse{}, err // Mengembalikan kesalahan jika promosi tidak ditemukan
	}
	existingPromotion.PromotionName = promo.PromotionName
	existingPromotion.DiscountType = promo.DiscountType
	existingPromotion.DiscountValue = promo.DiscountValue
	existingPromotion.PromotionStartDate = promo.PromotionStartDate
	existingPromotion.PromotionEndDate = promo.PromotionEndDate

	// Memperbarui promosi di repositori
	updatedPromotion, err := s.PromotionRepo.UpdatePromotionbyPromotionID(existingPromotion)
	if err != nil {
		return web.PromotionResponse{}, err // Mengembalikan kesalahan jika gagal memperbarui promosi
	}

	// Mengonversi promosi yang diperbarui menjadi respons yang sesuai
	response := helper.ToPromotionResponse(updatedPromotion)
	return response, nil
}

func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error{
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}

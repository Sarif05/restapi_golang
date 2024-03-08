package services

import (
	"errors"
	"final_project_promotion/internal/app/models"
	"final_project_promotion/internal/app/repositories"
	"final_project_promotion/utils/exceptions"

	"gorm.io/gorm"
)

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

func NewPromotionService (PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl{
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

func (s *PromotionServiceImpl) CreatePromotion (promo models.Promotion) (models.Promotion, error){
	return s.PromotionRepo.CreatePromotion(promo)
}

// func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error){
// 	return s.PromotionRepo.GetAllPromotions()
// }

func (s *PromotionServiceImpl) SearchPromotions(query string, limit, offset int) ([]models.Promotion, error) {
    return s.PromotionRepo.SearchPromotions(query, limit, offset)
}

func (s *PromotionServiceImpl) GetAllPromotions(limit, offset int) ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions(limit, offset)
}

func (s *PromotionServiceImpl) GetPromotionbyPromotionID (promotionID string) (models.Promotion, error){
	promo, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exceptions.PromotionIDNotFoundError{
				Message: "Promotion Not Found",
				PromotionID: promotionID,
			}
		}
		return models.Promotion{}, err
	}
	return promo, nil
}

func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID (promo models.Promotion) (models.Promotion, error){
	updatePromo, err := s.PromotionRepo.UpdatePromotionbyPromotionID(promo)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exceptions.PromotionIDNotFoundError{
				Message: "Duplicate Promotion Found",
				PromotionID: promo.PromotionID,
			}
		}
		return models.Promotion{}, err
	}
	return updatePromo, nil
}

func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error{
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}

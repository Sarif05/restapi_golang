package main

import (
	"final_project_promotion/internal/app/repositories"
	"final_project_promotion/internal/app/services"
	"final_project_promotion/internal/configs"
	"final_project_promotion/internal/delivery"

	"github.com/labstack/echo/v4"
)

func main(){

	configs.LoadViperEnv()

	db := configs.InitDatabase()

	e := echo.New()

	PromotionRepo := repositories.NewPromotionRepository(db)

	PromoService := services.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}

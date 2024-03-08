package delivery

import (
	"final_project_promotion/internal/app/handlers"
	"final_project_promotion/internal/app/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloServer(c echo.Context) error{
	return c.String(http.StatusOK, "Hello, World")
}

func PromotionRoute(e *echo.Echo, PromoService services.PromotionService){
	e.GET("/", HelloServer)
	// e.GET("/promotions", handlers.PSQLGetAllPromotionData(PromoService))
	e.GET("/promotions", handlers.PSQLGetAllPromotionData(PromoService))
	e.GET("/getpromotion/:promotion_id", handlers.PSQLGetPromotionbyPromotionID(PromoService))
	e.GET("/promotions/search", handlers.PSQLSearchPromotions(PromoService))
	e.POST("/createpromotion", handlers.PSQLCreatePromotionData(PromoService))
	e.PUT("/updatepromotion/:promotion_id", handlers.PSQLUpdatePromotionbyPromotionID(PromoService))
	e.DELETE("/deletepromotion/:promotion_id", handlers.PSQLDeletePromotionbyPromotionID(PromoService))
}
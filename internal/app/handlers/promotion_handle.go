package handlers

import (
	"final_project_promotion/internal/app/models"
	"final_project_promotion/internal/app/services"
	"final_project_promotion/utils/exceptions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PSQLCreatePromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	// Implementasi kamu taruh disini
	return func (c echo.Context) error {
		var promo models.Promotion
		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion Data")
		}

		createdPromo, err := PromoService.CreatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failde to create promotion")
		}
		return c.JSON(http.StatusCreated, createdPromo)
	}
}

// func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
// 	// Implementasi kamu taruh disini
// 	return func(c echo.Context) error {
// 		promotions, err := PromoService.GetAllPromotions()
// 		if err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions: " + err.Error())
// 		}
// 		return c.JSON(http.StatusOK, promotions)
// 	}
// }

func PSQLSearchPromotions(PromoService services.PromotionService) echo.HandlerFunc {
    return func(c echo.Context) error {
        query := c.QueryParam("query")
        limitParam := c.QueryParam("limit")
        pageParam := c.QueryParam("page")

        limit := 5  
        page := 1  

        if limitParam != "" {
            limit, _ = strconv.Atoi(limitParam)
        }

        if pageParam != "" {
            page, _ = strconv.Atoi(pageParam)
        }

        if limit < 1 {
            limit = 1
        }
        if page < 1 {
            page = 1
        }
        offset := (page - 1) * limit

        promotions, err := PromoService.SearchPromotions(query, limit, offset)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to search promotions: "+err.Error())
        }
        return c.JSON(http.StatusOK, promotions)
    }
}

func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Mendapatkan parameter opsional dari query string
        limitParam := c.QueryParam("limit")
        pageParam := c.QueryParam("page")

        // Default limit dan page jika tidak disediakan oleh pengguna
        limit := 5  // Jumlah data default
        page := 1   // Nomor halaman default

        // Konversi string ke integer untuk limit
        if limitParam != "" {
            limit, _ = strconv.Atoi(limitParam)
        }

        // Konversi string ke integer untuk page
        if pageParam != "" {
            page, _ = strconv.Atoi(pageParam)
        }

        if limit < 1 {
            limit = 1
        }
        if page < 1 {
            page = 1
        }
        offset := (page - 1) * limit
        promotions, err := PromoService.GetAllPromotions(limit, offset)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions: "+err.Error())
        }
        return c.JSON(http.StatusOK, promotions)
    }
}

func PSQLGetPromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	// Implementasi kamu taruh disini
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		promo, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil {

			// ! Update the exception with the custom one. For now leave it there.
			if e, ok := err.(*exceptions.PromotionIDNotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		return c.JSON(http.StatusOK, promo)
	}
}

func PSQLUpdatePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	// Implementasi kamu taruh disini
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		promo, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil{
			if e, ok := err.(*exceptions.PromotionIDNotFoundError); ok{
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failde To Get Promotion")
		}

		if err := c.Bind(&promo); err != nil{
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion Data")
		}
		updatePromo, err := PromoService.UpdatePromotionbyPromotionID(promo)
		if err != nil{
			return echo.NewHTTPError(http.StatusInternalServerError, "Failde To Update Promotion")
		}
		return c.JSON(http.StatusOK, updatePromo)
	}
}

func PSQLDeletePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	// Implementasi kamu taruh disini
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		if err := PromoService.DeletePromotionbyPromotionID(promotionID); err != nil {
			if e, ok := err.(*exceptions.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
		}
		return c.JSON(http.StatusNoContent, "Promotion Data deleted successfully") // 204
	}
}

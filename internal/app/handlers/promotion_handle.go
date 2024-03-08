package handlers

import (
	"final_project_promotion/internal/app/models/web"
	"final_project_promotion/internal/app/services"
	"final_project_promotion/utils/exceptions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func PSQLCreatePromotionData(s services.PromotionService) func(c echo.Context) error {
    return func(c echo.Context) error {
        promoCreate := web.PromotionCreateRequest{}
        if err := c.Bind(&promoCreate); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion Data")
        }
        createdPromo, err := s.CreatePromotion(promoCreate)
        if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}
        webResponse := web.WebResponse{
            Code:   http.StatusOK,
            Status: "OK",
            Data:   createdPromo,
        }
        return c.JSON(http.StatusOK, webResponse)
    }
}

func PSQLSearchPromotions(s services.PromotionService) func(c echo.Context) error {
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
        searchPromo, err := s.SearchPromotions(query, limit, offset)
        if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}
        webResponse := web.WebResponse{
            Code:   http.StatusOK,
            Status: "OK",
            Data:   searchPromo,
        }
        return c.JSON(http.StatusOK, webResponse)
    }
}

func PSQLGetAllPromotionData(s services.PromotionService) func(c echo.Context) error {
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
        getAll, err := s.GetAllPromotions(limit, offset)
        if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}
        webResponse := web.WebResponse{
            Code:   http.StatusOK,
            Status: "OK",
            Data:   getAll,
        }
        return c.JSON(http.StatusOK, webResponse)
    }
}

func PSQLGetPromotionbyPromotionID(s services.PromotionService) func(c echo.Context) error {
	// Implementasi kamu taruh disinildm
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		getById, err := s.GetPromotionbyPromotionID(promotionID)
        if err != nil {
			// ! Update the exception with the custom one. For now leave it there.
			if e, ok := err.(*exceptions.PromotionIDNotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}
        webResponse := web.WebResponse{
            Code:   http.StatusOK,
            Status: "OK",
            Data:   getById,
        }
		return c.JSON(http.StatusOK, webResponse)
	}
}

func PSQLUpdatePromotionbyPromotionID(PromoService services.PromotionService) func(c echo.Context) error {
    return func(c echo.Context) error {
        promotionID := c.Param("promotion_id")
        promoResponse, err := PromoService.GetPromotionbyPromotionID(promotionID)
        if err != nil {
			if e, ok := err.(*exceptions.PromotionIDNotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

        if err := c.Bind(&promoResponse); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

        // Bind data promosi yang diperbarui dari request
        updateRequest := web.PromotionUpdateRequest{}
        if err := c.Bind(&updateRequest); err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion Data")
        }
        // Set PromotionID dari updateRequest dengan nilai dari promoResponse
        updateRequest.PromotionID = promoResponse.PromotionID
        updatePromo, err := PromoService.UpdatePromotionbyPromotionID(updateRequest)
        if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update promotion")
		}

        // Buat respons sukses
        webResponse := web.WebResponse{
            Code:   http.StatusOK,
            Status: "OK",
            Data:   updatePromo,
        }
        return c.JSON(http.StatusOK, webResponse)
    }
}

func PSQLDeletePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	// Implementasi kamu taruh disini
	return func(c echo.Context) error {
		promotionID := c.Param("promotion_id")

		if err := PromoService.DeletePromotionbyPromotionID(promotionID); err != nil {
			if e, ok := err.(*exceptions.NotFoundErr); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
		}
		return c.JSON(http.StatusNoContent, "Promotion Data deleted successfully") // 204
	}
}

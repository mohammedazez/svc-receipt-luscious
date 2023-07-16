package handler

import (
	"net/http"
	port "svc-receipt-luscious/core/port/ingredient"
	"svc-receipt-luscious/interface/api/common"
	"svc-receipt-luscious/interface/api/extl/v1/ingredient/request"
	"svc-receipt-luscious/interface/api/extl/v1/ingredient/response"
	"svc-receipt-luscious/interface/api/utils/validation"

	domain "svc-receipt-luscious/core/domain/ingredient"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service port.Service
}

func NewHandler(service port.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) List(c echo.Context) error {
	req := new(request.RequestList)

	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ingredientList, err := h.service.List(req.IngredientName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := response.NewResponseList(http.StatusText(http.StatusOK), ingredientList, http.StatusOK)
	return c.JSON(http.StatusOK, res)

}

func (h *Handler) Insert(c echo.Context) error {

	req := new(request.RequestInsert)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := validation.ValidateReq(req)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, errVal)
	}

	ingredient := new(domain.IngredientService)
	ingredient.IngredientName = req.IngredientName
	ingredient.RecipeID = req.RecipeID
	ingredient.Quantity = req.Quantity

	err := h.service.Insert(ingredient)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusOK), http.StatusOK, true)
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) Update(c echo.Context) error {

	req := new(request.RequestUpdate)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := validation.ValidateReq(req)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, errVal)
	}

	ingredient := new(domain.IngredientService)
	ingredient.ID = c.Param("ingredient_id")
	ingredient.IngredientName = req.IngredientName
	ingredient.RecipeID = req.RecipeID
	ingredient.Quantity = req.Quantity

	err := h.service.Update(ingredient)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusOK), http.StatusOK, true)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Delete(c echo.Context) error {

	ingredientID := c.Param("ingredient_id")

	err := h.service.Delete(ingredientID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusOK), http.StatusOK, true)
	return c.JSON(http.StatusOK, res)
}

package handler

import (
	"net/http"
	port "svc-receipt-luscious/core/port/recipe"
	"svc-receipt-luscious/interface/api/common"
	"svc-receipt-luscious/interface/api/extl/v1/recipe/request"
	"svc-receipt-luscious/interface/api/extl/v1/recipe/response"
	"svc-receipt-luscious/interface/api/utils/validation"

	domain "svc-receipt-luscious/core/domain/recipe"

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

	categoryList, err := h.service.List(req.RecipeName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := response.NewResponseList(http.StatusText(http.StatusOK), categoryList, http.StatusOK)
	return c.JSON(http.StatusOK, res)
}

// func (h *Handler) Detail(c echo.Context) error {

// 	categoryID := c.Param("category_id")

// 	categoryList, err := h.service.Detail(categoryID)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}

// 	res := response.NewResponseDetail(http.StatusText(http.StatusOK), categoryList, http.StatusOK)
// 	return c.JSON(http.StatusOK, res)

// }

func (h *Handler) Insert(c echo.Context) error {

	req := new(request.RequestInsert)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	errVal := validation.ValidateReq(req)
	if errVal != nil {
		return c.JSON(http.StatusBadRequest, errVal)
	}

	category := new(domain.Recipe)
	category.RecipeName = req.RecipeName
	category.CategoryID = req.CategoryID
	category.HowToMake = req.HowToMake

	err := h.service.Insert(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusCreated), http.StatusCreated, true)
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

	recipe := new(domain.Recipe)
	recipe.ID = c.Param("recipe_id")
	recipe.RecipeName = req.RecipeName
	recipe.CategoryID = req.CategoryID
	recipe.HowToMake = req.HowToMake

	err := h.service.Update(recipe)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusOK), http.StatusOK, true)
	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Delete(c echo.Context) error {

	recipeID := c.Param("recipe_id")

	err := h.service.Delete(recipeID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := new(common.DefaultResponseNoData)
	res.SetResponseDataNoData(http.StatusText(http.StatusOK), http.StatusOK, true)
	return c.JSON(http.StatusOK, res)
}

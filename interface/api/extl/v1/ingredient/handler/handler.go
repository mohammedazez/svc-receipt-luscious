package handler

import (
	"net/http"
	port "svc-receipt-luscious/core/port/ingredient"
	"svc-receipt-luscious/interface/api/extl/v1/ingredient/request"
	"svc-receipt-luscious/interface/api/extl/v1/ingredient/response"
	"svc-receipt-luscious/utils/constant"

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

	res := response.NewResponseList(constant.SuccessGet, ingredientList)
	return c.JSON(http.StatusOK, res)

}

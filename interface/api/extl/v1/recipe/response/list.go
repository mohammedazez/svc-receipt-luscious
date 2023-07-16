package response

import (
	domain "svc-receipt-luscious/core/domain/recipe"
	"svc-receipt-luscious/interface/api/common"
)

type (
	ResponseList struct {
		Response
	}
)

func NewResponseList(message string, data []domain.Recipe, code int) *common.DefaultResponse {
	recipes := make([]ResponseList, 0)

	for _, val := range data {
		var recipe ResponseList

		recipe.ID = val.ID
		recipe.RecipeName = val.RecipeName
		recipe.CreatedAt = val.CreatedAt
		recipe.UpdatedAt = val.UpdatedAt
		recipes = append(recipes, recipe)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, recipes, code, true)
	return responseData
}

func NewResponseDetail(message string, category *domain.Recipe, code int) *common.DefaultResponse {
	data := new(Response)
	data.ID = category.ID
	data.CreatedAt = category.CreatedAt
	data.UpdatedAt = category.UpdatedAt

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, data, code, true)
	return responseData
}

package response

import (
	domain "svc-receipt-luscious/core/domain/ingredient"
	"svc-receipt-luscious/interface/api/common"
)

type (
	ResponseList struct {
		Response
	}
)

func NewResponseList(message string, data []domain.IngredientService, code int) *common.DefaultResponse {
	ingredients := make([]ResponseList, 0)

	for _, val := range data {
		var ingredient ResponseList

		ingredient.ID = val.ID
		ingredient.IngredientName = val.IngredientName
		ingredient.RecipeID = val.RecipeID
		ingredient.Quantity = val.Quantity
		ingredient.CreatedAt = val.CreatedAt
		ingredient.UpdatedAt = val.UpdatedAt
		ingredients = append(ingredients, ingredient)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, ingredients, code, true)
	return responseData
}

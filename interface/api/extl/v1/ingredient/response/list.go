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

func NewResponseList(message string, data []domain.IngredientService) *common.DefaultResponse {
	ingredients := make([]ResponseList, 0)

	for _, val := range data {
		var ingredient ResponseList

		ingredient.ID = val.ID
		ingredients = append(ingredients, ingredient)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, ingredients)
	return responseData
}

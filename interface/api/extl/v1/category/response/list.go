package response

import (
	domain "svc-receipt-luscious/core/domain/category"
	"svc-receipt-luscious/interface/api/common"
)

type (
	ResponseList struct {
		Response
	}
)

func NewResponseList(message string, data []domain.Category, code int) *common.DefaultResponse {
	categories := make([]ResponseList, 0)

	for _, val := range data {
		var category ResponseList

		category.ID = val.ID
		category.CategoryName = val.CategoryName
		category.CreatedAt = val.CreatedAt
		category.UpdatedAt = val.UpdatedAt
		categories = append(categories, category)
	}

	responseData := new(common.DefaultResponse)
	responseData.SetResponseData(message, categories, code, true)
	return responseData
}

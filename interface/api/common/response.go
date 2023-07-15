package common

type DefaultResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *DefaultResponse) SetResponseData(message string, data interface{}) {
	r.Message = message
	r.Data = data
}

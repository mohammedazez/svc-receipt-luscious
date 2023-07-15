package common

type DefaultResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func (r *DefaultResponse) SetResponseData(message string, data interface{}, code int, status bool) {
	r.Status = status
	r.Message = message
	r.Code = code
	r.Data = data
}

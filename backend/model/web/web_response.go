package web

type WebResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) WebResponse {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	responseWeb := WebResponse{
		Meta: meta,
		Data: data,
	}
	return responseWeb
}

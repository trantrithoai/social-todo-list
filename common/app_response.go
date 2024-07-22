package common

type SuccessResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data interface{}, paging interface{}, filter interface{}) *SuccessResponse {
	return &SuccessResponse{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{Data: data}
}

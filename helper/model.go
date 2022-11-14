package helper

// DefaultReponse: default struct of api response
type DefaultReponse struct {
	Message string `json:"message" example:"Bad Request"`
}

// DefaultDataResponse: default struct of api response with the data
type DefaultDataResponse struct {
	Data interface{} `json:"data"`
}

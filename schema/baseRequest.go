package schema

//Response is
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

//ResponseError is
type ResponseError struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

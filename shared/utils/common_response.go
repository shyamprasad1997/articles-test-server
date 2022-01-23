package utils

// CommonResponse responses common json data.
type CommonResponse struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

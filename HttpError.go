package restful

type HttpError struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

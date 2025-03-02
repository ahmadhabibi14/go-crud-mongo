package response

type ResponseCommon struct {
	StatusCode int    `json:"status"`            // HTTP Status Code (100 - 599)
	Message    string `json:"message,omitempty"` // Response message (if any)
	Error      string `json:"error,omitempty"`   // Error message (if any)
}

func (resp *ResponseCommon) SetStatus(code int) {
	resp.StatusCode = code
}

func (resp *ResponseCommon) SetMessage(msg string) {
	resp.Message = msg
}

func (resp *ResponseCommon) SetError(errStr string) {
	resp.Error = errStr
}

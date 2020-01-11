package apierrors

//APIError custom http error
type APIError struct {
	Code int    `json:"err_code"`
	Msg  string `json:"err_msg,omitempty"`
	Desc string `json:"err_desc,omitempty"`
}

//DefaulAPIError default error for success.
var DefaulAPIError = New(0, "Ok", "success")

//New new api error
func New(code int, msg, desc string) *APIError {
	return &APIError{
		Code: code,
		Msg:  msg,
		Desc: desc,
	}
}

//todo read from errors.yml. define errors in yaml file

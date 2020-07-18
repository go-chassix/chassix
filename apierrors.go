package chassis

import "c6x.io/chassis/apierrors"

//NewAPIError new api error
func NewAPIError(code int, msg, desc string) *apierrors.APIError {
	return apierrors.New(code, msg, desc)
}

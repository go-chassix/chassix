package chassis

import "pgxs.io/chassis/apierrors"

//NewAPIError new api error
func NewAPIError(code int, msg, desc string) *apierrors.APIError {
	return apierrors.New(code, msg, desc)
}

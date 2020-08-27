package chassix

import (
	"c6x.io/chassix.v2/apierrors"
)

//NewAPIError new api error
func NewAPIError(code int, msg, desc string) *apierrors.APIError {
	return apierrors.New(code, msg, desc)
}

package chassis

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

//Validate validate single instance
func Validate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}

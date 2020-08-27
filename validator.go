package chassix

import (
	"errors"
	"fmt"
	"sync"

	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslation "github.com/go-playground/validator/v10/translations/zh"

	"c6x.io/chassix.v2/apierrors"
)

var (
	validate          *validator.Validate
	validateOnce      sync.Once
	trans             ut.Translator
	errValidateFailed = errors.New("validate fialed")
)

//Validate validate single instance
func Validate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
		zh := zh.New()
		uni := ut.New(zh, zh)

		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		var found bool
		trans, found = uni.GetTranslator("zh")
		if !found {
			//todo check
		}

		zhTranslation.RegisterDefaultTranslations(validate, trans)
	})
	return validate
}

//ValidateTranslator validate trans
func ValidateTranslator() ut.Translator {
	Validate()
	return trans
}

//ValidateEntity validate entity
func ValidateEntity(entity interface{}) error {
	return Validate().Struct(entity)
}

//ValidateEntityAndWriteResp validate entity and write header and error as entity
func ValidateEntityAndWriteResp(res *restful.Response, entity interface{}, apiErr *apierrors.APIError) error {
	err := Validate().Struct(entity)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		customErr := *apiErr

		for _, e := range errs {
			// can translate each error one at a time.
			customErr.Desc = customErr.Desc + "[" + e.Translate(trans) + "]"
			fmt.Println(e.Translate(trans))
		}
		NewResponse(res).Error(400, &customErr)
		return errValidateFailed
	}
	return nil
}

package restful

import (
	"c6x.io/chassix.v2/config"
	restfulSpec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/go-openapi/spec"
)

// KeyOpenAPITags is a Metadata key for a restful Route
const KeyOpenAPITags = restfulSpec.KeyOpenAPITags

//newPostBuildOpenAPIObjectFunc open api api docs data
func newPostBuildOpenAPIObjectFunc() restfulSpec.PostBuildSwaggerObjectFunc {
	return func(swo *spec.Swagger) {
		swo.Info = &spec.Info{
			InfoProps: spec.InfoProps{
				Title:       config.Openapi().Spec.Title,
				Description: config.Openapi().Spec.Description,
				Contact: &spec.ContactInfo{
					Name:  config.Openapi().Spec.Contact.Name,
					Email: config.Openapi().Spec.Contact.Email,
					URL:   config.Openapi().Spec.Contact.URL,
				},

				License: &spec.License{
					Name: config.Openapi().Spec.License.Name,
					URL:  config.Openapi().Spec.License.URL,
				},
				Version: config.Openapi().Spec.Version,
			},
		}

		var nTags []spec.Tag
		for _, tag := range config.Openapi().Tags {
			nTag := spec.Tag{TagProps: spec.TagProps{Name: tag.Name, Description: tag.Description}}

			nTags = append(nTags, nTag)
		}
		swo.Tags = nTags
	}
}

//Serve rest webservice

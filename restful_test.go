package chassis

import (
	"reflect"
	"testing"

	"github.com/emicklei/go-restful/v3"
	"github.com/stretchr/testify/assert"
)

func Test_AddMetaDataTags(t *testing.T) {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/xxx").To(to))
	ws.Route(ws.POST("/xxxxx").To(to).Metadata(KeyOpenAPITags, []string{"test"}))
	tags := []string{"foo"}
	AddMetaDataTags(ws, tags)

	assert.NotEmpty(t, ws.Routes()[0].Metadata[KeyOpenAPITags])
	route2Tags, ok := ws.Routes()[1].Metadata[KeyOpenAPITags].([]string)
	assert.True(t, ok)
	assert.Equal(t, 2, len(route2Tags))
	assert.Equal(t, "test", route2Tags[0])
	assert.Equal(t, "foo", route2Tags[1])

}

func to(req *restful.Request, res *restful.Response) {
}

type TestEntity struct {
	Foo string
	Bar string
}

func Test_NewWriteSample(t *testing.T) {
	entity := NewWriteSample([]TestEntity{})
	assert.Equal(t, reflect.Slice, reflect.TypeOf(entity.Data).Kind())
}

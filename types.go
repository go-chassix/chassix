package chassix

import "github.com/go-chassix/chassix/v2/apierrors"

//BaseDTO Data Transfer Object
type BaseDTO struct {
	ID        uint     `json:"id,omitempty" description:"资源ID"`
	CreatedAt JSONTime `json:"created_at,omitempty" description:"创建日期"`
	UpdatedAt JSONTime `json:"updated_at,omitempty" description:"更新日期"`
}

//RespEntity common response entity
type RespEntity struct {
	apierrors.APIError
}

//EmptyRespEntity 空响应体
type EmptyRespEntity struct {
	apierrors.APIError
}

//PageDTO common page struct
type PageDTO struct {
	Total uint `json:"total"`
	Index uint `json:"page_index"`
	Size  uint `json:"page_size"`
	Pages uint `json:"pages"`
}

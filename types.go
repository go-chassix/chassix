package chassis

//BaseDTO Data Transfer Object
type BaseDTO struct {
	ID        uint     `json:"id,omitempty" description:"资源ID"`
	CreatedAt JSONTime `json:"created_at,omitempty" description:"创建日期"`
	UpdatedAt JSONTime `json:"updated_at,omitempty" description:"更新日期"`
}

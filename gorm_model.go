package chassis

import (
	"github.com/jinzhu/gorm"
	"time"
)

//Page page
type Page struct {
	List   interface{} `json:"list,omitempty"`
	Total  uint        `json:"total,omitempty"`
	Offset uint        `json:"offset,omitempty"`
	Index  uint        `json:"page_index,omitempty"`
	Size   uint        `json:"page_size,omitempty"`
	Pages  uint        `json:"pages,omitempty"`
}

//Pagination 新建分页查询
type Pagination struct {
	Offset    uint        `json:"offset,omitempty"`
	Limit     uint        `json:"limit,omitempty"`
	Condition interface{} `json:"condition,omitempty"`
}

//NewPage new page
func newPage(data interface{}, index, size, count uint) *Page {
	var pages uint
	if count%size == 0 {
		pages = count / size
	} else {
		pages = count/size + 1
	}
	return &Page{
		List:   data,
		Total:  count,
		Size:   size,
		Offset: index * size,
		Index:  index,
		Pages:  pages,
	}
}

//NewPagination pagination query
func NewPagination(db *gorm.DB, model interface{}, pageIndex, pageSize uint) *Page {
	var count uint
	db.Count(&count)
	if count > 0 && count > pageIndex*pageSize {
		db.Limit(pageSize).
			Offset(pageIndex * pageSize).
			Find(model)
		return newPage(model, pageIndex, pageSize, count)
	}
	return nil
}

//SampleModel model with id pk
type SampleModel struct {
	ID uint `gorm:"primary_key" json:"id"`
}

//Model gorm model
type Model struct {
	ID        uint      `gorm:"primary_key" json:"id"`            // primary key
	CreatedAt JSONTime  `json:"created_at,omitempty"`             // created time
	UpdatedAt JSONTime  `json:"updated_at,omitempty"`             //updated time
	DeletedAt *JSONTime `sql:"index" json:"deleted_at,omitempty"` //deleted time
	Version   uint      `json:"version,omitempty"`                //version opt lock
}

//ComplexModel gorm model composed Model add Addtion
type ComplexModel struct {
	Model
	Addition string `json:"addition,omitempty"`
}

// JSONTime 格式化输出时间
type JSONTime time.Time

// MarshalJSON json解码
func (j JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Local().Format("2006-01-02 15:04:05") + `"`), nil
}

// JSONDate 格式化输出时间
type JSONDate time.Time

// MarshalJSON json解码
func (j JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Local().Format("2006-01-02") + `"`), nil
}

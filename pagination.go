package chassix

const (
	//SortDirectionDESC 降序
	SortDirectionDESC = "DESC"
	//SortDirectionASC 升序
	SortDirectionASC = "ASC"
)

//Pageable 可分页参数
type Pageable struct {
	Page  uint
	Size  uint
	Sorts []Sort
}

//Sort  分页查询排序
type Sort struct {
	Field     string
	Direction string
}

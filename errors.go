package chassis

/**
* Common errors. code prefix 100xxx
 */

//ErrIDInvalid REST resource ID invalid
var ErrIDInvalid = NewAPIError(100001, "非法ID", "ID为数值且不能为0")

//ErrPageParamsInvalid Pagination params invalid
var ErrPageParamsInvalid = NewAPIError(100002, "分页参数有误", "page_index,page_size为数值类型,page_size 大于5小于100 被5整除")

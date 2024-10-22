package serializer

// ===============================================================================
// = 通用ID对应的DTO
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// ===============================================================================
// = 通用的query
type CommonQueryOtpones struct {
	Filter map[string]interface{} // 永久过滤器格式 {"desc":"Hello, World!","CreatedAt":1728866342120,"title":"fas"}
	Ranges CommonPage             // 分页  [0, 99]
	Sort   CommonSort             // 排序  sort  ["desc","ASC"]
}

// ===============================================================================
// = 通用的分页

type CommonPage struct {
	Skip  int
	Limit int
}

// ===============================================================================
// = 排序
type CommonSort struct {
	Field string
	Md    string
}

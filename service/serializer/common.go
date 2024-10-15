package serializer

// ===============================================================================
// = 通用ID对应的DTO
type CommonIDDTO struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// ===============================================================================
// = 通用的query
type CommonQueryOtpones struct {
	Filter interface{} // 永久过滤器格式 {"desc":"Hello, World!","CreatedAt":1728866342120,"title":"fas"}
	Ranges []any       // 分页  [0, 99]
	Sort   []any       // 排序  sort  ["desc","ASC"]
}

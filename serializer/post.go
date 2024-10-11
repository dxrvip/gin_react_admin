package serializer

// 序列化器
type PostRequry struct {
	Title   string `json:"title" binding:"required,min=4,max=100"`
	Content string `json:"content" binding:"required,min=4,max=10000"`
	Cid     int    `json:"cid" binding:"required,min=1"`
	Desc    string `json:"desc" binding:"omitempty,max=255"`
	Img     string `json:"img" binding:"omitempty,url"`
}

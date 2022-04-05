package dateMessage

// type DateBetween struct {
// 	CreateTime time.Time `form:"create_time" binding:"required,timing" time_format:"2006-01-02" label:"起始日期"`
// 	UpdateTime time.Time `form:"update_time" binding:"required,timing" time_format:"2006-01-02" label:"更新日期"`
// }
type DateBetween struct {
	CreateTime string `form:"create_time" json:"create_time,omitempty" binding:"timing" time_format:"2006-01-02" label:"起始日期"`
	UpdateTime string `form:"update_time" json:"update_time,omitempty" binding:"timing" time_format:"2006-01-02" label:"更新日期"`
}

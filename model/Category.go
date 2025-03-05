package model

// 文章
type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`                         // ID Gorm默认以ID为主键
	Name string `json:"name" gorm:"type:varchar(20);not null;unique;"` //分类名称	类型为varchar(20)	不为空	唯一
	//CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`              // 创建时间
	//UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`              // 必须为UpdatedAt   UpdateAt错误          // 更新时间
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"` // 创建时间
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"` // 必须为UpdatedAt   UpdateAt错误          // 更新时间

}

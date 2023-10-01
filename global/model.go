package global

import (
	"time"
)

type GVA_MODEL struct {
	Id        uint      `gorm:"primarykey"` // 主键ID
	CreateTime time.Time `gorm:"createTime" json:"createTime"` // 创建时间

}

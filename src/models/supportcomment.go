package models

import (
	"time"
)

type Supportcomment struct {
	Id             int64     `json:"Id" xorm:"pk autoincr comment('自增ID') BIGINT(20)"`
	Commentid      int64     `json:"CommentId" xorm:"not null default 0 comment('评论标识') unique(supportsatellite_I2) BIGINT(20)"`
	Useridentifier int64     `json:"UserIdentifier" xorm:"not null default 0 comment('用户标识') unique(supportsatellite_I2) BIGINT(20)"`
	Status         int       `json:"Status" xorm:"not null default 0 comment('0代表禁用，1代表启用。') TINYINT(4)"`
	Createtime     time.Time `json:"CreateTime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	Updatetime     time.Time `json:"UpdateTime" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Appid          int64     `json:"AppId" xorm:"not null default 11 comment('应用标识') unique(supportsatellite_I2) BIGINT(20)"`
}

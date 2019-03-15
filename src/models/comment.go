package models

import (
	"time"
)

type Comment struct {
	Id             int64     `json:"id,omitempty" xorm:"pk autoincr comment('消息ID') BIGINT(20)"`
	Content        string    `json:"content,omitempty" xorm:"not null comment('消息内容') index TEXT"`
	Fatherid       int64     `json:"-" xorm:"not null default 0 comment('消息父级ID') index(comment_I14) index(comment_I15) index(comment_I16) index(comment_I17) index(comment_I18) index index(comment_I8) index(comment_i2) BIGINT(255)"`
	Images         string    `json:"images,omitempty" xorm:"not null default '' comment('消息图片地址') VARCHAR(1000)"`
	Ssid           int64     `json:"-" xorm:"not null default 0 comment('信息类型主体标识') index(comment_I10) index(comment_I11) index(comment_I12) index(comment_I13) index(comment_I14) index(comment_I15) index(comment_I16) index(comment_I17) index(comment_I18) index(comment_I7) index(comment_I8) index(comment_I9) index(comment_i2) BIGINT(20)"`
	Title          string    `json:"-" xorm:"not null default '' comment('信息标题') VARCHAR(100)"`
	Url            string    `json:"url,omitempty" xorm:"not null default '' comment('信息主体地址') VARCHAR(200)"`
	Ip             string    `json:"-" xorm:"not null default '' comment('用户IP地址') VARCHAR(20)"`
	Place          string    `json:"-" xorm:"not null default '' comment('用户ip所在位置') VARCHAR(20)"`
	Supportcount   int64     `json:"support_count,omitempty" xorm:"not null default 0 comment('支持数') index(comment_I11) index(comment_I8) index(comment_i2) BIGINT(255)"`
	Iselite        int       `json:"-" xorm:"not null default 0 comment('是否加精') index(comment_I10) index(comment_I16) index(comment_I17) TINYINT(255)"`
	Istop          int       `json:"-" xorm:"not null default 0 comment('是否置顶') index(comment_I16) index(comment_I17) index(comment_I7) index(comment_I8) index(comment_I9) index(comment_i2) TINYINT(255)"`
	Status         int       `json:"-" xorm:"not null default 0 comment('信息状态，0为删除，1为正常') TINYINT(255)"`
	Revertcount    int64     `json:"revert_count,omitempty" xorm:"not null default 0 comment('回复数') BIGINT(255)"`
	Useridentifier int64     `json:"-" xorm:"not null default 0 comment('用户标识') index(comment_i3) BIGINT(20)"`
	Appid          int64     `json:"-" xorm:"not null default 0 comment('应用程序标识') index(comment_I15) index(comment_i2) BIGINT(20)"`
	Createtime     time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') index(comment_I12) index(comment_I13) index(comment_I14) index(comment_I15) index(comment_I17) index(comment_I18) index index(comment_I7) index(comment_i3) DATETIME"`
	Updatetime     time.Time `json:"-" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Ssidtype       int       `json:"-" xorm:"not null default 0 comment('信息主体标识，0代表漫画评论，1代表社区评论') index(comment_I7) TINYINT(255)"`
	Relateid       string    `json:"-" xorm:"not null default '' comment('关连Id') index(comment_I13) index(comment_I14) VARCHAR(20)"`
	Contentlength  int64     `json:"-" xorm:"not null default 0 comment('评论内容长度') index(comment_I17) BIGINT(20)"`
	Iswater        int       `json:"-" xorm:"not null default 0 comment('是否水贴，0代表不是，1代表是') index(comment_I16) index(comment_I17) TINYINT(255)"`
	Haveimage      int       `json:"have_image,omitempty" xorm:"not null default 0 comment('是否有图片') TINYINT(4)"`
	//
	CreatetimeOut int64 `json:"create_time,omitempty"`
	IsSupport     int   `json:"is_support,omitempty"`
}

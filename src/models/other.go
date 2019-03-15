package models

type WholeComment struct {
	UserInfoItem *UserInfo `json:"user_info,omitempty"`
	CommentItem  *Comment  `json:"comment,omitempty"`
}

type UserInfo struct {
	Uid    int64  `json:"uid,omitempty" xorm:"not null pk INT(11)"`
	Uname  string `json:"name,omitempty" xorm:"not null VARCHAR(20)"`
	Ulevel int    `json:"level,omitempty" xorm:"not null TINYINT(4)"`
	IsVip  int    `json:"is_vip,omitempty"`
}

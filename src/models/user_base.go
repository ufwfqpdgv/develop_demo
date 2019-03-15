package models

type UserBase struct {
	Channel       string `json:"channel" xorm:"not null VARCHAR(50)"`
	KmhUid        int    `json:"kmh_uid" xorm:"not null INT(11)"`
	Platformid    int    `json:"platformid" xorm:"not null INT(11)"`
	Productlineid int    `json:"productlineid" xorm:"not null INT(11)"`
	Siteid        int    `json:"siteid" xorm:"not null INT(11)"`
	Uareacode     string `json:"Uareacode" xorm:"not null VARCHAR(50)"`
	UbindDevice   string `json:"Ubind_device" xorm:"not null VARCHAR(50)"`
	UbindHuawei   string `json:"Ubind_huawei" xorm:"not null VARCHAR(50)"`
	UbindKmh      string `json:"Ubind_kmh" xorm:"not null VARCHAR(50)"`
	UbindMkxq     string `json:"Ubind_mkxq" xorm:"not null VARCHAR(50)"`
	UbindQq       string `json:"Ubind_qq" xorm:"not null VARCHAR(50)"`
	UbindQqunion  string `json:"Ubind_qqunion" xorm:"not null VARCHAR(50)"`
	UbindSina     string `json:"Ubind_sina" xorm:"not null VARCHAR(50)"`
	UbindWeixin   string `json:"Ubind_weixin" xorm:"not null VARCHAR(50)"`
	Ubindinfo     string `json:"Ubindinfo" xorm:"not null VARCHAR(250)"`
	Ubirthday     string `json:"Ubirthday" xorm:"not null VARCHAR(20)"`
	Uid           int    `json:"Uid" xorm:"not null pk INT(11)"`
	Uisguest      int    `json:"Uisguest" xorm:"not null VARCHAR(50)"`
	Ulevel        int    `json:"Ulevel" xorm:"not null TINYINT(4)"`
	Umail         string `json:"Umail" xorm:"not null VARCHAR(50)"`
	Uname         string `json:"Uname" xorm:"not null VARCHAR(20)"`
	Uregip        string `json:"Uregip" xorm:"not null VARCHAR(15)"`
	Uregtime      string `json:"Uregtime" xorm:"not null VARCHAR(20)"`
	Usex          int    `json:"Usex" xorm:"not null TINYINT(4)"`
	Usign         string `json:"Usign" xorm:"not null VARCHAR(50)"`
	Ustatus       int    `json:"Ustatus" xorm:"not null TINYINT(4)"`
	UtagIds       string `json:"Utag_ids" xorm:"not null VARCHAR(100)"`
	Utype         int    `json:"Utype" xorm:"not null TINYINT(4)"`
	Uviptime      string `json:"Uviptime" xorm:"not null VARCHAR(20)"`
	//
	IsVip int `json:"isVip"`
}

CREATE database IF NOT EXISTS samh_vip;

DROP TABLE IF EXISTS samh_vip.user_vip;
CREATE TABLE samh_vip.user_vip (
  user_vip_id BIGINT NOT NULL AUTO_INCREMENT,
  uid BIGINT NOT NULL DEFAULT 0 COMMENT '用户id',
  gold_expire_time BIGINT DEFAULT 0 COMMENT '黄金会员到期时间，当同一用户拥有多种会员时，如加钻石会员时间时得将黄金会员的时候后延至钻石会员结束的时间点',
  diamond_expire_time BIGINT DEFAULT 0 COMMENT '钻石会员到期时间，当同一用户拥有多种会员时，如加钻石会员时间时得将黄金会员的时候后延至钻石会员结束的时间点',
  create_time BIGINT DEFAULT 0,
  update_time BIGINT DEFAULT 0,
  status INT DEFAULT 0 COMMENT '0-正常，1-已删除',
  PRIMARY KEY (user_vip_id),
  UNIQUE KEY (uid)
) DEFAULT charset=utf8mb4 COMMENT '用户会员信息表';

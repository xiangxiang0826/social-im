CREATE TABLE `app_limit`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `cur_day_num` int(11) NOT NULL DEFAULT '0' COMMENT '当前日期天数字',
    `uid`         bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
    `cur_nums`    int(11) NOT NULL DEFAULT '0' COMMENT '当前计数',
    `limit_type`  tinyint(3) NOT NULL DEFAULT '0' COMMENT '限制类型(1:派对,2:家族)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `type_uid_day_num` (`uid`,`cur_day_num`,`limit_type`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `app_room_mic`
(
    `id`                   BIGINT (20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `mark`                 VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '房间唯一标识',
    `name`                 VARCHAR(60)  NOT NULL DEFAULT '' COMMENT '房间名称',
    `background_url`       VARCHAR(650) NOT NULL DEFAULT '' COMMENT '房间背景图片',
    `background_small_url` VARCHAR(650) NOT NULL DEFAULT '' COMMENT '房间背景小图',
    `party_type`           TINYINT (2) NOT NULL DEFAULT '0' COMMENT '派对类型',
    `create_user`          BIGINT (20) NOT NULL DEFAULT '0' COMMENT '房间创建人',
    `online_nums`          INT (10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '在线用户数',
    `created_at`           DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`           DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `mark` (`mark`),
    KEY `create_user` (`create_user`)
) ENGINE = INNODB AUTO_INCREMENT = 0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT = '派对房';
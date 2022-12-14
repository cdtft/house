-- house.house definition

CREATE TABLE `house`
(
    `id`          int(11)                                 NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `unit`        varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '单元号',
    `floor_num`   varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '楼层',
    `house_no`    varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '房号',
    `idx`         int(11)                                          DEFAULT NULL COMMENT '顺序编码',
    `floorage`    varchar(20) COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '建筑面积',
    `sold`        tinyint(1)                              NOT NULL DEFAULT '0' COMMENT '是否售卖',
    `create_time` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '创建时间',
    `task_id`     int(11)                                          DEFAULT NULL COMMENT '任务号',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 500
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- house.task definition

CREATE TABLE `task`
(
    `id`        int(11)      NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `task_time` varchar(100) NOT NULL DEFAULT '' COMMENT '任务时间',
    `total`     int(11)               DEFAULT NULL COMMENT '卖出数量',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4;
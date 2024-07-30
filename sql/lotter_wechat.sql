drop table if exists  `t_prize`;
create table `t_prize`
(
     `id` int(10) unsigned not null auto_increment,
    `name` varchar(255) not null default '' comment '奖品名称',
    `pic` varchar(255) not null default '' comment '奖品图片',
    `link` varchar(255) not null default '' comment '奖品链接',
    `type` int(10) unsigned not null default '0'  comment '奖品类型，1-虚拟币，2-虚拟券，3-实物小奖，4-实物大奖',
    `data` varchar(255) not null default '' comment '奖品数据',
    `total` int(11) not null default '-1' comment '奖品数量，0 无限量 》0限量 《0无奖品',
    `left` int(11) not null default '0' comment '剩余数量',
    `is_use` int(10) unsigned not null default '0' comment '是否使用中 1-使用中 2-未使用',
    `probability` int(11) not null default '0' comment'中奖概率，万分之',
    `probability_max` int(11) not null default '0' comment '中奖概率上限',
    `probability_min` int(11) not null default '0' comment'中奖概率下限',
    primary key(`id`)
)comment ='奖品表';
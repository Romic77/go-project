create table `book`(
                       `id` int(11) auto_increment NOT NULL,
                       `title` varchar(255) NOT NULL,
                       `price` double NOT NULL,
                       PRIMARY key(`id`)
)engine=innoDB auto_increment=1 default charset=utf8;
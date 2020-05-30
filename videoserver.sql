create table comments(
    id varchar (64) not null,
    video_id varchar (64),
    author_id int(10),
    content text,
    time datetime default  current_timestamp
);

create table sessions(
    session_id varchar (64) not null ,
    TTL varchar (64),
    user_name text
);
alter table sessions add primary key (session_id);

create table users(
    id int unsigned not null auto_increment,
    user_name varchar(64),
    pwd text not null,
    unique key (user_name),
    primary key (id)
);

create table video_del_rec(
    id int unsigned not null auto_increment,
    video_id varchar(64),
    primary key (id)
);

create table video_info(
    id varchar(64) not null ,
    author_id int(10),
    name varchar(64),
    display_ctime text,
    create_time datetime default current_timestamp,
    primary key (id)
)
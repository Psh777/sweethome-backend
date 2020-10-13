create table sensors
(
    id               uuid                                      not null
        constraint sensors_pk_2
            primary key,
    comment          text        default ''::text              not null,
    room             varchar(50) default ''::character varying not null,
    enable           boolean     default true                  not null,
    update_timestamp timestamp                                 not null,
    alive            integer     default 0                     not null,
    request_id       uuid
);

create index sensors_room_index
    on sensors (room);

create table sensors_data
(
    sensor_id    uuid                         not null,
    sensor_type  integer        default 0     not null,
    sensor_value numeric(10, 2) default 0     not null,
    timestamp    timestamp      default now() not null,
    id           bigserial                    not null
        constraint "sensors-data_pk"
            primary key,
    request_id   uuid
);

create table logs
(
    time timestamp     default now()                 not null,
    log  varchar(5000) default ''::character varying not null
);

create table switch
(
    id   varchar(100) default ''::character varying not null
        constraint switch_pk
            primary key,
    room varchar(50)  default ''::character varying not null,
    ip   varchar(20)  default ''::character varying not null,
    port integer      default 0                     not null
);

create unique index switch_id_uindex
    on switch (id);

create table telegram_chats
(
    id integer default 0 not null
);


create unique index telegram_chats_id_uindex
    on telegram_chats (id);

create table devices
(
    id          varchar(100) default ''::character varying not null
        constraint devices_pk
            primary key,
    type        varchar(50)  default ''::character varying,
    alisa_type  varchar(50)  default ''::character varying,
    room        varchar(50)  default ''::character varying not null,
    name        varchar(50)  default ''::character varying not null,
    description varchar(200) default ''::character varying not null,
    url         varchar(200) default ''::character varying not null
);

create unique index devices_id_uindex
    on devices (id);

create table capabilities
(
    id        serial                                     not null
        constraint capabilities_pk
            primary key,
    device_id varchar(50)  default ''::character varying not null,
    type      varchar(200) default ''::character varying not null,
    state     varchar(50)  default ''::character varying not null,
    instance  varchar      default ''::character varying not null
);


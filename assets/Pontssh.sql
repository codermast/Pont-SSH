create table main.sshs
(
    id       TEXT
        primary key,
    host     TEXT    not null,
    port     INTEGER not null,
    username TEXT    not null,
    password TEXT    not null,
    comment  TEXT
);

create table main.logs
(
    id        text
        constraint logs_pk
            primary key,
    server_id text,
    command   text,
    time      text,
    exec_date date
);


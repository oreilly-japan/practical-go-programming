create table author (
    id   integer PRIMARY KEY,
    name varchar(99) not null,
    created_at timestamp not null default now()
);

create table book (
    id        integer PRIMARY KEY,
    title     varchar(99) not null,
    price     integer not null,
    author_id integer not null,
    created_at timestamp not null default now()
);

alter table book add foreign key (author_id) references author (id);

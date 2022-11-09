CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    password_hash varchar(255) not null,
    date_creation timestamp    not null,
    last_update   timestamp    not null
);

CREATE TABLE categories
(
    id            serial                                                 not null unique,
    user_id       int           references users (id) on delete cascade  not null,
    name          varchar(255)                                           not null,
    date_creation timestamp                                              not null,
    last_update   timestamp                                              not null
);

CREATE TABLE words
(
    id            serial                                                     not null unique,
    user_id       int           references users (id) on delete cascade      not null,
    category_id   int           references categories (id) on delete cascade not null,
    name          varchar(255)                                               not null,
    date_creation timestamp                                                  not null,
    last_update   timestamp                                                  not null
);

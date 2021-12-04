CREATE TABLE users
(
    id           serializable not null unique,
    username     varchar(255) not null unique,
    name         varchar(255) not null,
    passwordHash varchar(255) not null
);

CREATE TABLE todo_lists
(
    id          serializable not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE user_lists
(
    id serializable                                      not null unique,
    user_id references users (id) on delete cascade      not null,
    list_id references todo_lists (id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id          serializable not null unique,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);

CREATE TABLE list_items
(
    id serializable                                      not null unique,
    list_id references todo_lists (id) on delete cascade not null,
    item_id references todo_items (id) on delete cascade not null
);


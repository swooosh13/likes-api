-- +goose Up
-- +goose StatementBegin
create table container_item
(
    id           serial
        constraint container_item_pk
            primary key,
    container_id int
        constraint container_item_container_id_fk
            references container,
    name         varchar,
    symbol       varchar,
    priority int default 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table container_item;
-- +goose StatementEnd

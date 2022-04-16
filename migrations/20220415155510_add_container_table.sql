-- +goose Up
-- +goose StatementBegin
create table container
(
    id      serial
        constraint container_pk
            primary key,
    user_id varchar not null,
    name    varchar(40)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table container;
-- +goose StatementEnd

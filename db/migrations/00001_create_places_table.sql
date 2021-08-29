-- +goose Up
-- +goose StatementBegin
CREATE table places(
                       id serial primary key,
                       memo varchar(255) not null,
                       seat varchar(255) not null,
                       user_id int not null,
                       created_at timestamptz not null default now(),
                       updated_at timestamptz
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table places
-- +goose StatementEnd

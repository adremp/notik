-- +goose Up
-- +goose StatementBegin
alter table pages add column created_at timestamp not null default now();
alter table pages add column updated_at timestamp not null default now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table pages drop column created_at;
alter table pages drop column updated_at;
-- +goose StatementEnd

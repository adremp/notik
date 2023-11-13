-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
    id serial primary key,
    title VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pages;
-- +goose StatementEnd

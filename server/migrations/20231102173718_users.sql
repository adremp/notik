-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
	id serial primary key,
	username VARCHAR(80) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	email VARCHAR(80) NOT NULL,
	password VARCHAR(80) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd

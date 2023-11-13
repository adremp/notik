-- +goose Up
-- +goose StatementBegin
CREATE TYPE part_variant AS ENUM ('header1', 'header2', 'header3', 'text', 'image');
CREATE TABLE parts (
    id serial primary key,
    body TEXT NOT NULL,
		part_order int NOT NULL,
    variant part_variant NOT NULL, 
    page_id BIGINT NOT NULL,
    FOREIGN KEY (page_id) REFERENCES pages(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE parts;
DROP TYPE part_variant;
-- +goose StatementEnd

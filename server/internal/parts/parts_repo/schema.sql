CREATE TYPE part_type AS ENUM ('header1', 'header2', 'header3', 'text', 'image');

CREATE TABLE parts (
	id serial PRIMARY KEY,
	part_order smallint not null,
	variant part_type NOT NULL,
	title VARCHAR(255) NOT NULL,
	page_id BIGINT NOT NULL REFERENCES pages(id)
);
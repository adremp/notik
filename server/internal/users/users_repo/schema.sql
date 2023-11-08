CREATE TABLE users (
	id serial primary key,
	username VARCHAR(80) NOT NULL,
	email VARCHAR(80) NOT NULL,
	password VARCHAR(80) NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
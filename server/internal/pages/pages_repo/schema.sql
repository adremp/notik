CREATE TABLE pages (
	id serial PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	user_id BIGINT NOT NULL REFERENCES users(id)
);
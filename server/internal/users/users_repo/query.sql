-- name: Create :one
insert into
	users (username, email, password)
values
	($1, $2, $3) RETURNING id, username, email;


-- name: GetByFields :many
select * from users 
where (id = COALESCE(NULLIF(@id::int, 0), id)) AND 
(username = COALESCE(NULLIF(@username::text, ''), username)) AND 
(email = COALESCE(NULLIF(@email::text, ''), email)) 
limit COALESCE(NULLIF(@limits::int, 0), 1);


-- name: Delete :exec
delete from
	users
where
	id = $1;



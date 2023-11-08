-- name: Create :one
insert into
	users (username, email, password)
values
	($1, $2, $3) RETURNING id, username, email;


-- name: GetByEmail :one 
select username, email, password from users where email = $1;

-- name: Delete :exec
delete from
	users
where
	id = $1;

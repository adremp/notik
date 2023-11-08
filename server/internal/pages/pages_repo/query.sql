

-- name: GetById :one
select
	*
from
	pages p
where
	p.id = $1;


-- name: Create :one
insert into pages (
	title,
	user_id
) values (
	$1,
	$2
) returning *;
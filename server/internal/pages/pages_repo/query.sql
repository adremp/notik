

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
) returning id, title;



-- name: GetByFields :many
select id, title, user_id from pages 
where (id = COALESCE(NULLIF(@id::int, 0), id)) AND 
(title = COALESCE(NULLIF(@title::text, ''), title)) AND 
(user_id = COALESCE(NULLIF(@user_id::text, ''), user_id)) 
limit COALESCE(NULLIF(@limits::int, 0), 1) 
offset @offsets::int;
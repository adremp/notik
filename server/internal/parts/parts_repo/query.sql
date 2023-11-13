-- name: Upsert :one
insert into parts (variant, body, page_id) values ($1, $2, $3) 
on conflict (id) do 
update set variant = $1, body = $2, page_id = $3 returning variant, body, page_id, id;



-- name: GetByFields :many
select * from parts 
where (id = COALESCE(NULLIF(@id::int, 0), id)) AND 
(username = COALESCE(NULLIF(@username::text, ''), username)) AND 
(email = COALESCE(NULLIF(@email::text, ''), email)) 
limit COALESCE(NULLIF(@limits::int, 0), 1);
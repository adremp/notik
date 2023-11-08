-- name: Create :one
with part_el as (
	select
		coalesce(max(part_order), 0) + 1 as part_order
	from
		parts
	where
		page_id = $3
)
insert into
	parts (part_order, variant, title, page_id)
values
	(part_el.part_order, $1, $2, $3) RETURNING *;
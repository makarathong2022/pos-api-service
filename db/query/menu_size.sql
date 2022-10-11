-- name: CreateMenuSize :one
INSERT INTO tbl_menu_sizes (
    size_cd,
    size_name,
    sort,
    description
  )
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMenuSize :exec 
UPDATE tbl_menu_sizes SET
    size_cd = $2,
    size_name = $3,
    sort = $4,
    description = $5
    WHERE id = $1;

-- name: DeleteMenuSize :exec
UPDATE tbl_menu_sizes SET 
    deleted_at = $2
    WHERE id = $1;

-- name: GetMenuSizes :many
SELECT * FROM tbl_menu_sizes ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetMenuSize :one
SELECT * FROM tbl_menu_sizes WHERE id = $1 LIMIT 1;




-- name: CreateMenuGroup :one
INSERT INTO tbl_menu_groups (
    group_cd,
    group_name,
    sort,
    description
  )
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMenuGroup :exec
UPDATE tbl_menu_groups SET
    group_cd = $2,
    group_name = $3,
    sort = $4,
    description = $5
    WHERE id = $1;

-- name: DeleteMenuGroup :exec
UPDATE tbl_menu_groups SET deleted_at = $2 WHERE id = $1;

-- name: GetMenuGroup :one
SELECT * FROM tbl_menu_groups WHERE id = $1;

-- name: GetMenuGroups :many
SELECT * FROM tbl_menu_groups ORDER BY created_at LIMIT $1 OFFSET $2;


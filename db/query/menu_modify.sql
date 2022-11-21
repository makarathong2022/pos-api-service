-- name: CreateMenuModify :one
INSERT INTO tbl_menu_modifies(modify_cd, modify_name, sort, description) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMenuModify :exec
UPDATE tbl_menu_modifies SET 
    modify_cd = $2,
    modify_name = $3,
    sort = $4,
    description = $5,
    updated_at = DATE.NOW()
    WHERE id = $1 RETURNING *;

-- name: DeleteMenuModify :exec
UPDATE tbl_menu_modifies SET 
    deleted_at = DATE.NOW()
    WHERE id = $1 RETURNING *;

-- name: GetMenuModifies :many
SELECT * FROM tbl_menu_modifies ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetMenuModify :one
SELECT * FROM tbl_menu_modifies WHERE id = $1 LIMIT 1;
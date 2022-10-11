-- name: CreateMenuCategory :one
INSERT INTO tbl_menu_categories(category_cd, category_name, sort, description)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateMenuCategory :exec
UPDATE tbl_menu_categories SET category_cd  = $2, category_name = $3, sort = $4, description = $5, updated_at = $6
WHERE id = $1;

-- name: GetMenuCategories :many
SELECT * FROM tbl_menu_categories ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetMenuCategory :one
SELECT * FROM tbl_menu_categories WHERE id = $1 LIMIT 1;

-- name: DeleteMenuCategory :exec
UPDATE tbl_menu_categories SET deleted_at = DATE.NOW() WHERE id = $1;
-- name: CreateMenuItem :one
INSERT INTO tbl_menu_items (
    item_cd,
    item_name,
    abv,
    sort,
    outlet_id,
    description
  )
VALUES (
    $1, $2, $3, $4, $5, $6 
)RETURNING *;


-- name: GetMenuItem :one
SELECT * FROM tbl_menu_items
WHERE id = $1 LIMIT 1;

-- name: GetMenuItems :many
SELECT * FROM tbl_menu_items
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdateMenuItem :exec
UPDATE tbl_menu_items SET 
 item_cd = $2,
 item_name = $3, 
 abv = $4, 
 sort = $5, 
 outlet_id = $6,
 description = $7,
 updated_at = $8 
 WHERE id = $1 RETURNING *;

 -- name: DeleteMenuItem :exec 
 UPDATE tbl_menu_items SET
 deleted_at = date.now() 
 WHERE id = $1 RETURNING *;
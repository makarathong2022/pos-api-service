-- name: CreateMenuItemModify :one

INSERT INTO tbl_menu_item_modifies (item_cd, modify_cd) VALUES ($1, $2) RETURNING *;

-- name: UpdateMenuItemModify :exec

UPDATE tbl_menu_item_modifies SET item_cd = $2, modify_cd = $3 WHERE id = $1 RETURNING *;

-- name: GetMenuItemModifies :many

SELECT m.modify_cd, im.modify_name, m.item_cd, i.item_name FROM tbl_menu_item_modifies m 
    INNER JOIN tbl_menu_modifies im ON m.modify_cd = im.modify_cd
    INNER JOIN tbl_menu_items i ON i.item_cd = m.item_cd   
    LIMIT $1 OFFSET $2;

-- name: GetMenuItemModify :one
SELECT * FROM tbl_menu_item_modifies WHERE id = $1 LIMIT 1;

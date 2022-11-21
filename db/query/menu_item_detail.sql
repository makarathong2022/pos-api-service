-- name: CreateMenuItemDetail :one
INSERT INTO tbl_menu_item_details (
    item_cd_detail,
    item_id,
    group_id,
    category_id,
    size_id,
    cost,
    price,
    vat_id,
    vat,
    terminal_id,
    outlet_id,
    printer_id
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *;


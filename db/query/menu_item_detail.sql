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


-- name: GetMenuItemDetials :many
SELECT d.item_cd_detail, d.item_id, i.item_name, gp.group_cd, gp.group_name,
         c.category_cd, c.category_name, s.size_cd, s.size_name,
         d.cost, d.price, p.sort, v.vat_cd, v.vat_name, p.print_cd,
         p.print_name, p.ip_address, d.created_at, d.updated_at,
         d.deleted_at
         FROM tbl_menu_item_details d 
         INNER JOIN tbl_menu_items i ON d.item_id = i.item_cd 
         INNER JOIN tbl_menu_categories c ON c.category_cd = d.category_id
         INNER JOIN tbl_menu_sizes s ON s.size_cd = d.size_id
         INNER JOIN tbl_menu_groups gp ON gp.group_cd = d.group_id
         INNER JOIN tbl_outlets o ON o.outlet_cd = d.outlet_id
         INNER JOIN tbl_printers p ON p.print_cd = d.printer_id
         INNER JOIN tbl_vats v ON v.vat_cd = d.vat_id
         LIMIT $1 OFFSET $2;

-- name: GetMenuItemDetial :one
SELECT d.item_cd_detail, d.item_id, i.item_name, gp.group_cd, gp.group_name,
         c.category_cd, c.category_name, s.size_cd, s.size_name,
         d.cost, d.price, p.sort, v.vat_cd, v.vat_name, p.print_cd,
         p.print_name, p.ip_address, d.created_at, d.updated_at,
         d.deleted_at
         FROM tbl_menu_item_details d 
         INNER JOIN tbl_menu_items i ON d.item_id = i.item_cd 
         INNER JOIN tbl_menu_categories c ON c.category_cd = d.category_id
         INNER JOIN tbl_menu_sizes s ON s.size_cd = d.size_id
         INNER JOIN tbl_menu_groups gp ON gp.group_cd = d.group_id
         INNER JOIN tbl_outlets o ON o.outlet_cd = d.outlet_id
         INNER JOIN tbl_printers p ON p.print_cd = d.printer_id
         INNER JOIN tbl_vats v ON v.vat_cd = d.vat_id
         WHERE d.id = $1
         LIMIT 1;


-- name: UpdateMenuItemDetail :exec
UPDATE tbl_menu_item_details SET 
       item_id = $2, group_id = $3, category_id = $4,
       size_id = $5, vat_id = $6, cost = $7, price = $8,
       outlet_id = $9, printer_id  = $10, terminal_id = $11,
       updated_at = DATE.NOW()
       WHERE id = $1;

-- name: DeteleMenuItemDetail :exec
UPDATE tbl_menu_item_details SET deleted_at = DATE.NOW() WHERE id = $1;
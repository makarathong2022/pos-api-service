-- name: GetTaxs :many
SELECT * FROM tbl_vats LIMIT $1 OFFSET $2;

-- name: GetTax :one
SELECT * FROM tbl_vats WHERE ID = $1 LIMIT 1;

-- name: CreateTax :one
INSERT INTO tbl_vats (
    vat_cd,
    vat_key,
    vat_name,
    sort,
    description
  )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateTax :exec
UPDATE tbl_vats SET vat_cd = $2, vat_key = $3, vat_name = $4, sort = $5, description = $6, updated_at = DATE.NOW() WHERE id = $1 RETURNING *;

-- name: DeleteTax :exec
UPDATE tbl_vats SET deleted_at = DATE.NOW() WHERE id = $1 RETURNING *;
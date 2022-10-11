-- name: CreateOutlet :one
INSERT INTO tbl_outlets (
    outlet_cd,
    outlet_name,
    ip_address,
    sort,
    description
  )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdateOutlet :exec
UPDATE tbl_outlets SET
    outlet_cd = $2,
    outlet_name = $3,
    ip_address = $4,
    sort = $5,
    description = $6,
    updated_at = $7
    WHERE id = $1 RETURNING *;

-- name: DeleteOutlet :exec
UPDATE tbl_outlets SET
    deleted_at = $2
    WHERE id = $1 RETURNING *;

-- name: GetOutlets :many
SELECT * FROM tbl_outlets ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetOutlet :one
SELECT * FROM tbl_outlets WHERE id = $1;
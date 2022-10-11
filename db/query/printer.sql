-- name: CreatePrinter :one
INSERT INTO tbl_printers (
    print_cd,
    print_name,
    sort,
    ip_address,
    description
  )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: UpdatePrinter :exec
UPDATE tbl_printers SET
    print_cd = $2,
    print_name = $3,
    sort = $4,
    ip_address = $5,
    description = $6
    WHERE id = $1;

-- name: DeletePrinter :exec
UPDATE tbl_printers SET
    deleted_at = $2
    WHERE id = $1;

-- name: GetPrinters :many
SELECT * FROM tbl_printers LIMIT $1 OFFSET $2;

-- name: GetPrinter :one
SELECT * FROM tbl_printers WHERE id = $1;
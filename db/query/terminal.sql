-- name: CreateTerminal :one
INSERT INTO tbl_terminals (
    terminal_cd,
    terminal_name,
    ip_address,
    sort,
    outlet_cd,
    description
  )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: UpdateTerminal :exec
UPDATE tbl_terminals SET 
    terminal_cd = $2,
    terminal_name = $3,
    ip_address = $4,
    sort = $5,
    outlet_cd = $6,
    description = $7,
    updated_at = $8
    WHERE id = $1; 

-- name: DeleteTerminal :exec
UPDATE tbl_terminals SET
    deleted_at = $2
    WHERE id = $1;

-- name: GetTerminals :many
SELECT * FROM tbl_terminals ORDER BY created_at LIMIT $1 OFFSET $2;

-- name: GetTerminal :one
SELECT * FROM tbl_terminals WHERE id = $1;
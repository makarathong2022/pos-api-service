// Code generated by sqlc. DO NOT EDIT.
// source: menu_group.sql

package db

import (
	"context"
	"time"
)

const createMenuGroup = `-- name: CreateMenuGroup :one
INSERT INTO tbl_menu_groups (
    group_cd,
    group_name,
    sort,
    description
  )
VALUES ($1, $2, $3, $4) RETURNING id, group_cd, group_name, sort, description, created_at, updated_at, deleted_at
`

type CreateMenuGroupParams struct {
	GroupCd     int32  `json:"group_cd"`
	GroupName   string `json:"group_name"`
	Sort        int64  `json:"sort"`
	Description string `json:"description"`
}

func (q *Queries) CreateMenuGroup(ctx context.Context, arg CreateMenuGroupParams) (TblMenuGroup, error) {
	row := q.db.QueryRowContext(ctx, createMenuGroup,
		arg.GroupCd,
		arg.GroupName,
		arg.Sort,
		arg.Description,
	)
	var i TblMenuGroup
	err := row.Scan(
		&i.ID,
		&i.GroupCd,
		&i.GroupName,
		&i.Sort,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const deleteMenuGroup = `-- name: DeleteMenuGroup :exec
UPDATE tbl_menu_groups SET deleted_at = $2 WHERE id = $1
`

type DeleteMenuGroupParams struct {
	ID        int64     `json:"id"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (q *Queries) DeleteMenuGroup(ctx context.Context, arg DeleteMenuGroupParams) error {
	_, err := q.db.ExecContext(ctx, deleteMenuGroup, arg.ID, arg.DeletedAt)
	return err
}

const getMenuGroup = `-- name: GetMenuGroup :one
SELECT id, group_cd, group_name, sort, description, created_at, updated_at, deleted_at FROM tbl_menu_groups WHERE id = $1
`

func (q *Queries) GetMenuGroup(ctx context.Context, id int64) (TblMenuGroup, error) {
	row := q.db.QueryRowContext(ctx, getMenuGroup, id)
	var i TblMenuGroup
	err := row.Scan(
		&i.ID,
		&i.GroupCd,
		&i.GroupName,
		&i.Sort,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getMenuGroups = `-- name: GetMenuGroups :many
SELECT id, group_cd, group_name, sort, description, created_at, updated_at, deleted_at FROM tbl_menu_groups ORDER BY created_at LIMIT $1 OFFSET $2
`

type GetMenuGroupsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetMenuGroups(ctx context.Context, arg GetMenuGroupsParams) ([]TblMenuGroup, error) {
	rows, err := q.db.QueryContext(ctx, getMenuGroups, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TblMenuGroup{}
	for rows.Next() {
		var i TblMenuGroup
		if err := rows.Scan(
			&i.ID,
			&i.GroupCd,
			&i.GroupName,
			&i.Sort,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMenuGroup = `-- name: UpdateMenuGroup :exec
UPDATE tbl_menu_groups SET
    group_cd = $2,
    group_name = $3,
    sort = $4,
    description = $5
    WHERE id = $1
`

type UpdateMenuGroupParams struct {
	ID          int64  `json:"id"`
	GroupCd     int32  `json:"group_cd"`
	GroupName   string `json:"group_name"`
	Sort        int64  `json:"sort"`
	Description string `json:"description"`
}

func (q *Queries) UpdateMenuGroup(ctx context.Context, arg UpdateMenuGroupParams) error {
	_, err := q.db.ExecContext(ctx, updateMenuGroup,
		arg.ID,
		arg.GroupCd,
		arg.GroupName,
		arg.Sort,
		arg.Description,
	)
	return err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: companies.sql

package database

import (
	"context"
)

const getCompany = `-- name: GetCompany :one
SELECT id, name, email, created_at, updated_at FROM company where id = $1
`

func (q *Queries) GetCompany(ctx context.Context, id int32) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompany, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getCompanyPackages = `-- name: GetCompanyPackages :many
SELECT id, name, company_id, price, terms, created_at, updated_at FROM package where company_id = $1
`

func (q *Queries) GetCompanyPackages(ctx context.Context, companyID int32) ([]Package, error) {
	rows, err := q.db.QueryContext(ctx, getCompanyPackages, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Package
	for rows.Next() {
		var i Package
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CompanyID,
			&i.Price,
			&i.Terms,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listCompanies = `-- name: ListCompanies :many
SELECT id, name, email, created_at, updated_at FROM company
`

func (q *Queries) ListCompanies(ctx context.Context) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, listCompanies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Company
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.CreatedAt,
			&i.UpdatedAt,
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

-- name: ListCompanies :many
SELECT * FROM company;

-- name: GetCompany :one
SELECT * FROM company where id = $1;

-- name: GetCompanyPackages :many
SELECT * FROM package where company_id = $1;
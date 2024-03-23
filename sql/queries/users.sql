-- name: GetUser :one
SELECT * FROM policyholder where telnumber = $1;

-- name: GetUserCompantServices :many
SELECT * FROM service 
JOIN package on service.company_id = package.company_id
JOIN subscription ON package.id = subscription.package_id 
where policyholder_id = $1;

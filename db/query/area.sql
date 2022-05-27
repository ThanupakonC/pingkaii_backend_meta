-- name: ListProvinces :many
SELECT *
FROM province
ORDER BY province_id;

-- name: ListProvincesDistrict :many
SELECT *
FROM district
WHERE province_id = $1
ORDER BY district_id;

-- name: ListDistrictTambon :many
SELECT *
FROM tambon
WHERE district_id = $1
ORDER BY tambon_id;
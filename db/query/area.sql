-- name: ListProvinces :many
SELECT *
FROM province
ORDER BY province_thai;

-- name: ListProvincesDistrict :many
SELECT *
FROM district
WHERE province_id = $1
ORDER BY district_thai;

-- name: ListDistrictTambon :many
SELECT *
FROM tambon
WHERE district_id = $1
ORDER BY tambon_thai;

-- name: ListTambonPostcode :many
SELECT *
FROM postcode
WHERE tambon_id = $1
ORDER BY postcode;
-- name: GetStockRatingByTicker :one
SELECT id, ticker, company, brokerage, action, rating_from, rating_to,
       target_from, target_to, created_at, upside, change_target, current_price
FROM challenge.stock_rating
WHERE ticker = $1;

-- name: UpsertStockRating :exec
INSERT INTO challenge.stock_rating (
    ticker, company, brokerage, action, rating_from, rating_to,
    target_from, target_to, upside, change_target, current_price
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
ON CONFLICT (ticker) DO UPDATE SET
    company = EXCLUDED.company,
    brokerage = EXCLUDED.brokerage,
    action = EXCLUDED.action,
    rating_from = EXCLUDED.rating_from,
    rating_to = EXCLUDED.rating_to,
    target_from = EXCLUDED.target_from,
    target_to = EXCLUDED.target_to,
    current_price = EXCLUDED.current_price,
    upside = EXCLUDED.upside,
    change_target = EXCLUDED.change_target;

-- name: ListStockRatings :many
SELECT 
    id, ticker, company, brokerage, action, rating_from, rating_to,
    target_from, target_to, created_at, upside, change_target, current_price
FROM challenge.stock_rating
WHERE 
($1::text = '' OR 
        (ticker ILIKE '%' || $1::text || '%') OR
        (company ILIKE '%' || $1::text || '%') OR
        (brokerage IS NOT NULL AND brokerage ILIKE '%' || $1::text || '%') OR
        (action IS NOT NULL AND action ILIKE '%' || $1::text || '%') OR
        (rating_from IS NOT NULL AND rating_from ILIKE '%' || $1::text || '%') OR
        (rating_to IS NOT NULL AND rating_to ILIKE '%' || $1::text || '%'))
AND ($6::numeric IS NULL OR $6::numeric = 0 OR upside >= $6)
AND ($7::numeric IS NULL OR $7::numeric = 0 OR current_price >= $7)
AND ($8::numeric IS NULL OR $8::numeric = 0 OR current_price <= $8)
ORDER BY 
    CASE WHEN $2::text = 'ticker' AND $3::text = 'asc' THEN ticker END ASC,
    CASE WHEN $2::text = 'ticker' AND $3::text = 'desc' THEN ticker END DESC,
    CASE WHEN $2::text = 'company' AND $3::text = 'asc' THEN company END ASC,
    CASE WHEN $2::text = 'company' AND $3::text = 'desc' THEN company END DESC,
    CASE WHEN $2::text = 'rating_from' AND $3::text = 'asc' THEN rating_from END ASC,
    CASE WHEN $2::text = 'rating_from' AND $3::text = 'desc' THEN rating_from END DESC,
    CASE WHEN $2::text = 'rating_to' AND $3::text = 'asc' THEN rating_to END ASC,
    CASE WHEN $2::text = 'rating_to' AND $3::text = 'desc' THEN rating_to END DESC,
    CASE WHEN $2::text = 'target_from' AND $3::text = 'asc' THEN target_from END ASC,
    CASE WHEN $2::text = 'target_from' AND $3::text = 'desc' THEN target_from END DESC,
    CASE WHEN $2::text = 'target_to' AND $3::text = 'asc' THEN target_to END ASC,
    CASE WHEN $2::text = 'target_to' AND $3::text = 'desc' THEN target_to END DESC,
    CASE WHEN $2::text = 'created_at' AND $3::text = 'asc' THEN created_at END ASC,
    CASE WHEN $2::text = 'created_at' AND $3::text = 'desc' THEN created_at END DESC,
    CASE WHEN $2::text = 'upside' AND $3::text = 'asc' THEN upside END ASC,
    CASE WHEN $2::text = 'upside' AND $3::text = 'desc' THEN upside END DESC,
    CASE WHEN $2::text = 'change_target' AND $3::text = 'asc' THEN change_target END ASC,
    CASE WHEN $2::text = 'change_target' AND $3::text = 'desc' THEN change_target END DESC,
    CASE WHEN $2::text = 'current_price' AND $3::text = 'asc' THEN current_price END ASC,
    CASE WHEN $2::text = 'current_price' AND $3::text = 'desc' THEN current_price END DESC,
    id ASC
LIMIT $4::int
OFFSET $5::int;

-- name: CountStockRatings :one
SELECT COUNT(*)
FROM challenge.stock_rating
WHERE 
    ($1::text = '' OR 
            (ticker IS NOT NULL AND ticker ILIKE '%' || $1::text || '%') OR
            (company IS NOT NULL AND company ILIKE '%' || $1::text || '%') OR
            (brokerage IS NOT NULL AND brokerage ILIKE '%' || $1::text || '%') OR
            (action IS NOT NULL AND action ILIKE '%' || $1::text || '%') OR
            (rating_from IS NOT NULL AND rating_from ILIKE '%' || $1::text || '%') OR
            (rating_to IS NOT NULL AND rating_to ILIKE '%' || $1::text || '%'))
AND ($2::numeric IS NULL OR $2::numeric = 0 OR upside >= $2)
AND ($3::numeric IS NULL OR $3::numeric = 0 OR current_price >= $3)
AND ($4::numeric IS NULL OR $4::numeric = 0 OR current_price <= $4);


-- name: InsertStockRatingHistory :exec
INSERT INTO challenge.stock_rating_history (
    stock_rating_id,
    date,
    old_current_price,
    new_current_price,
    old_upside,
    new_upside
) VALUES ($1, NOW(), $2, $3, $4, $5);

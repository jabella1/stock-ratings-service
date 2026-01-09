-- name: GetStockRatingByTicker :one
SELECT id, ticker, company, brokerage, action, rating_from, rating_to,
       target_from, target_to, created_at
FROM challenge.stock_rating
WHERE ticker = $1;

-- name: UpsertStockRating :exec
INSERT INTO challenge.stock_rating (
    ticker, company, brokerage, action, rating_from, rating_to,
    target_from, target_to
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
)
ON CONFLICT (ticker) DO UPDATE SET
    company = EXCLUDED.company,
    brokerage = EXCLUDED.brokerage,
    action = EXCLUDED.action,
    rating_from = EXCLUDED.rating_from,
    rating_to = EXCLUDED.rating_to,
    target_from = EXCLUDED.target_from,
    target_to = EXCLUDED.target_to;

-- name: ListStockRatings :many
SELECT 
    id, ticker, company, brokerage, action, rating_from, rating_to,
    target_from, target_to, created_at
FROM challenge.stock_rating
WHERE 
($1::text = '' OR 
        (ticker ILIKE '%' || $1::text || '%') OR
        (company ILIKE '%' || $1::text || '%') OR
        (brokerage IS NOT NULL AND brokerage ILIKE '%' || $1::text || '%') OR
        (action IS NOT NULL AND action ILIKE '%' || $1::text || '%') OR
        (rating_from IS NOT NULL AND rating_from ILIKE '%' || $1::text || '%') OR
        (rating_to IS NOT NULL AND rating_to ILIKE '%' || $1::text || '%'))
ORDER BY 
    CASE WHEN $2::text = 'ticker' AND $3::text = 'asc' THEN ticker END ASC,
    CASE WHEN $2::text = 'ticker' AND $3::text = 'desc' THEN ticker END DESC,
    CASE WHEN $2::text = 'company' AND $3::text = 'asc' THEN company END ASC,
    CASE WHEN $2::text = 'company' AND $3::text = 'desc' THEN company END DESC,
    CASE WHEN $2::text = 'target_from' AND $3::text = 'asc' THEN target_from END ASC,
    CASE WHEN $2::text = 'target_from' AND $3::text = 'desc' THEN target_from END DESC,
    CASE WHEN $2::text = 'target_to' AND $3::text = 'asc' THEN target_to END ASC,
    CASE WHEN $2::text = 'target_to' AND $3::text = 'desc' THEN target_to END DESC,
    CASE WHEN $2::text = 'created_at' AND $3::text = 'asc' THEN created_at END ASC,
    CASE WHEN $2::text = 'created_at' AND $3::text = 'desc' THEN created_at END DESC,
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
            (rating_to IS NOT NULL AND rating_to ILIKE '%' || $1::text || '%'));
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
-- name: GetStockRatingByTicker :one
SELECT id, ticker, company, brokerage, action, rating_from, rating_to,
       target_from, target_to, created_at
FROM challenge.stock_rating
WHERE ticker = $1;
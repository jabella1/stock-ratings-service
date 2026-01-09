ALTER TABLE challenge.stock_rating 
ADD CONSTRAINT stock_rating_ticker_key UNIQUE (ticker);
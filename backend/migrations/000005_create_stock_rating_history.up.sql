CREATE TABLE challenge.stock_rating_history (
    id BIGINT NOT NULL DEFAULT unique_rowid(),
    stock_rating_id BIGINT NOT NULL,
    date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    old_current_price DECIMAL(10,2) NULL,
    new_current_price DECIMAL(10,2) NOT NULL,
    old_upside DECIMAL(6,2) NULL,
    new_upside DECIMAL(6,2) NOT NULL,
    CONSTRAINT stock_rating_history_pkey PRIMARY KEY (id),
    CONSTRAINT fk_stock_rating FOREIGN KEY (stock_rating_id)
        REFERENCES challenge.stock_rating(id)
        ON DELETE CASCADE
);

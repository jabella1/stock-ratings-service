CREATE SCHEMA IF NOT EXISTS challenge;

CREATE TABLE challenge.stock_rating (
  id INT8 NOT NULL DEFAULT unique_rowid(),
  ticker STRING NOT NULL,
  company STRING NOT NULL,
  brokerage STRING NULL,
  action STRING NULL,
  rating_from STRING NULL,
  rating_to STRING NULL,
  target_from DECIMAL(10,4) NULL,
  target_to DECIMAL(10,4) NULL,
  created_at TIMESTAMPTZ NULL DEFAULT NOW(),
  CONSTRAINT stock_rating_pkey PRIMARY KEY (id)
);
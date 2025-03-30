-- Down Migration: Drop indexes first, then drop the reviews table

-- Drop indexes
DROP INDEX IF EXISTS idx_reviews_user_id;
DROP INDEX IF EXISTS idx_reviews_rating;

-- Drop table
DROP TABLE IF EXISTS reviews;

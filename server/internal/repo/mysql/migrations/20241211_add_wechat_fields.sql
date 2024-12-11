-- Add WeChat-related columns to user table
ALTER TABLE "user" 
ADD COLUMN IF NOT EXISTS open_id VARCHAR(255),
ADD COLUMN IF NOT EXISTS union_id VARCHAR(255),
ADD COLUMN IF NOT EXISTS nickname VARCHAR(255),
ADD UNIQUE (open_id);

-- Migration UP
ALTER TABLE contas
ALTER COLUMN senha TYPE VARCHAR(255);

ALTER TABLE IF EXISTS projects ADD COLUMN network varchar(256);
UPDATE projects SET network = 'MAINNET';
ALTER TABLE IF EXISTS projects ALTER COLUMN network SET NOT NULL;
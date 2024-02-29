-- Create the 'credentials' table with an index on the 'id' column
CREATE TABLE IF NOT EXISTS credentials (
  id UUID PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  website VARCHAR(255) NOT NULL,
  category VARCHAR(255),
  owner_id UUID NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  health SMALLINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
  FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE
);

-- Add an index to the 'id' column
CREATE INDEX credential_idx_id ON credentials (id);

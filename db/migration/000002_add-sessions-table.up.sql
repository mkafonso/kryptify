-- Create the 'sessions' table with an index on the 'id' column
CREATE TABLE IF NOT EXISTS sessions (
  id UUID PRIMARY KEY,
  account_id VARCHAR(255) NOT NULL,
  refresh_token TEXT NOT NULL,
  user_agent TEXT,
  client_ip VARCHAR(255),
  is_blocked BOOLEAN,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT (now())
);

-- Add an index to the 'id' column
CREATE INDEX session_idx_id ON sessions (id);

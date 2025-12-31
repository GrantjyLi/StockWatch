CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE watchlists (
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES users(id),
  name TEXT,
  created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE alerts (
  id UUID PRIMARY KEY,
  watchlist_id UUID REFERENCES watchlists(id),
  ticker TEXT NOT NULL,
  operator TEXT CHECK (operator IN ('>=', '<=', '=')),
  target_price NUMERIC NOT NULL,

  triggered BOOLEAN DEFAULT false,
  triggered_at TIMESTAMP,

  created_at TIMESTAMP DEFAULT now()
);

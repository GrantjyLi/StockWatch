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

CREATE INDEX index_watchlists_user_id ON watchlists(user_id);
CREATE INDEX index_alerts_watchlist_id ON alerts(watchlist_id);
CREATE INDEX index_alerts_poller_main ON alerts (
    ticker,
    triggered,
    created_at
);
import React from "react";

import { useState } from "react";
import Header from "./components/Header";
import WatchlistGrid from "./components/WatchlistGrid";

export default function App() {
  const [watchlists, setWatchlists] = useState([
    {
      id: 1,
      name: "Watchlist #1",
      enabled: true,
      items: [
        { symbol: "AAPL", condition: ">=", price: 1190 },
        { symbol: "AAPL", condition: "<=", price: 1000 },
        { symbol: "AAPL", condition: "=", price: 1300 },
      ],
    },
    {
      id: 2,
      name: "Watchlist #2",
      enabled: true,
      items: [
        { symbol: "AAPL", condition: ">=", price: 1190 },
        { symbol: "AAPL", condition: "<=", price: 1000 },
        { symbol: "AAPL", condition: "=", price: 1300 },
      ],
    },
  ]);

  const addWatchlist = () => {
    setWatchlists([
      ...watchlists,
      {
        id: Date.now(),
        name: `Watchlist #${watchlists.length + 1}`,
        enabled: true,
        items: [],
      },
    ]);
  };

  const toggleWatchlist = (id) => {
    setWatchlists((prev) =>
      prev.map((w) =>
        w.id === id ? { ...w, enabled: !w.enabled } : w
      )
    );
  };

  return (
    <>
      <Header />
      <WatchlistGrid
        watchlists={watchlists}
        onAdd={addWatchlist}
        onToggle={toggleWatchlist}
      />
    </>
  );
}

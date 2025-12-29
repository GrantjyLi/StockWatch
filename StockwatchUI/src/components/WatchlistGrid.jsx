import React from "react";


import WatchlistCard from "./WatchlistCard";
import WatchlistCardNew from "./WatchlistCardNew";

export default function WatchlistGrid({ watchlists, onAdd, onToggle }) {
  return (
    <div style={styles.container}>
      {watchlists.map((w) => (
        <WatchlistCard
          key={w.id}
          watchlist={w}
          onToggle={() => onToggle(w.id)}
        />
      ))}
      <WatchlistCardNew onAdd={onAdd} />
    </div>
  );
}

const styles = {
  container: {
    background: "#B9B7B7",
    margin: 20,
    padding: 30,
    borderRadius: 40,
    display: "flex",
    gap: 20,
    flexWrap: "wrap",
  },
};

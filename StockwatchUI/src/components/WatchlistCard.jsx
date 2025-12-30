import React from "react";

export default function WatchlistCard({ watchlistData, onToggle }) {
  return (
    <div style={styles.card}>
      <h3>{watchlistData.name}</h3>
      <small>{watchlistData.items.length} Items</small>

      <div style={styles.items}>
        {watchlistData.items.map((i, idx) => (
          <div key={idx} style={styles.item}>
            <span>{i.symbol}</span>
            <span>
              {i.condition}${i.price}
            </span>
          </div>
        ))}
        {watchlistData.items.length === 0 && <span>…</span>}
      </div>

      <label style={styles.toggle}>
        <span>{watchlistData.enabled ? "On" : "Off"}</span>
        <input
          type="checkbox"
          checked={watchlist.enabled}
          onChange={onToggle}
        />
      </label>
    </div>
  );
}

const styles = {
    card: {
        background: "#A6FFA6",
        borderRadius: 20,
        padding: 20,
        width: 250,
        fontFamily: "Arial, Helvetica, sans-serif"
    },
    items: {
        background: "rgba(0,0,0,0.1)",
        borderRadius: 10,
        padding: 10,
        marginTop: 10,
    },
    item: {
        display: "flex",
        justifyContent: "space-between",
    },
    toggle: {
        marginTop: 10,
        display: "flex",
        justifyContent: "space-between",
    },
};

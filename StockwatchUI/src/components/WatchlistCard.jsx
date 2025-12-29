import React from "react";


export default function WatchlistCard({ watchlist, onToggle }) {
  return (
    <div style={styles.card}>
      <h3>{watchlist.name}</h3>
      <small>{watchlist.items.length} Items</small>

      <div style={styles.items}>
        {watchlist.items.map((i, idx) => (
          <div key={idx} style={styles.item}>
            <span>{i.symbol}</span>
            <span>
              {i.condition}${i.price}
            </span>
          </div>
        ))}
        {watchlist.items.length === 0 && <span>…</span>}
      </div>

      <label style={styles.toggle}>
        <span>{watchlist.enabled ? "On" : "Off"}</span>
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

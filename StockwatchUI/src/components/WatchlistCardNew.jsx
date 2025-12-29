import React from "react";


export default function WatchlistCardNew({ onAdd }) {
  return (
    <div style={styles.card} onClick={onAdd}>
      <h3>New Watchlist</h3>
      <div style={styles.plus}>+</div>
    </div>
  );
}

const styles = {
  card: {
    background: "#fff",
    borderRadius: 20,
    padding: 20,
    width: 250,
    cursor: "pointer",
    textAlign: "center",
    fontFamily: "Arial, Helvetica, sans-serif"
  },
  plus: {
    fontSize: 40,
  },
};

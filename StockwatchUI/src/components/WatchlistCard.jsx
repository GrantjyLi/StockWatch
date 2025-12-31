import React from "react";

export default function WatchlistCard({ WID, watchlistData }) {
    const tickers = watchlistData.tickers

    return (
        <div style={styles.card}>
            <h3>{watchlistData.name}</h3>
            <small>{tickers.length} Items</small>

            <div style={styles.items}>
            {Object.entries(tickers).map(([ticker, condition]) => (
                <div key={WID} style={styles.item}>
                <span>{ticker}</span>
                <span>
                    {condition}
                </span>
                </div>
            ))}
            {tickers.length === 0 && <span>…</span>}
            </div>
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

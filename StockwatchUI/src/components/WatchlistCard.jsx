import React from "react";

export default function WatchlistCard({ WID, watchlistData }) {
    const alerts = watchlistData.alerts
    console.log(alerts)
    return (
        <div style={styles.card}>
            <h3>{watchlistData.name}</h3>
            <small>{alerts.length} Items</small>

            <div style={styles.items}>
            {Object.entries(alerts).map(([alertID, ticker, operator, price]) => (
                <div key={alertID} style={styles.item}>
                <span>{ticker}</span>
                <span>
                    {operator}{price}
                </span>
                </div>
            ))}
            {alerts.length === 0 && <span>…</span>}
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

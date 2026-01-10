import React from "react";

export default function WatchlistCard({ WID, watchlistData, deleteWatchlist }) {
    const alerts = watchlistData.alerts
    
    return (
        <div style={styles.card} key={WID}>
            <button
                style={styles.deleteBtn}
                onClick={()=>{deleteWatchlist(WID, watchlistData.name)}}
            >
                ✕
            </button>
            <h3>{watchlistData.name}</h3>
            <small>{alerts.length} Items</small>

            <div style={styles.items}>
            {alerts.map((alert) => (
                <div key={alert.ID} style={styles.item}>
                <span>{alert.ticker}</span>
                <span>
                    {`${alert.operator} ${alert.price}`}
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
        fontFamily: "Arial, Helvetica, sans-serif",
        position: "relative"
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
    deleteBtn: {
        position: "absolute",
        top: 8,
        right: 8,
        padding: "4px 8px",
        fontSize: 12,
        cursor: "pointer",
        fontWeight: "bold",        
        background: "transparent",  
        border: "none",             
  },
};

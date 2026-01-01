import React from "react";

import WatchlistCard from "./WatchlistCard";
import WatchlistCardNew from "./WatchlistCardNew";

export default function WatchlistGrid({ watchlists, onAdd, onDelete }) {
    return (
    <div style={styles.container}>
        {Object.entries(watchlists).map(([id, data]) => (
        <WatchlistCard
            WID={id}
            watchlistData={data}
            deleteWatchlist={onDelete}
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

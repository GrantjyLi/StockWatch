import React from "react";

export default function WatchlistAddAlert ({ alertData, onChange, onRemove }) {
    return (
        <div style={styles.body}>

        <input style={styles.input} placeholder="Ticker"
            onChange={(e) => onChange({ ...alertData, ticker: e.target.value })}
        />
        <select style={styles.input} name="cars" id="cars"
            onChange={(e) => onChange({ ...alertData, operator: e.target.value })}
        >
            <option value=">=">&gt;=</option>
            <option value="<=">&lt;=</option>
            <option value="=">=</option>
        </select>
        <input style={styles.input} placeholder="Price" type="number" step="any" min="0" name="amount"
            onChange={(e) => onChange({ ...alertData, price: e.target.value })}
        />
        <button style={styles.removeX} onClick={onRemove}>✕</button>
    </div>
    )
}

const styles = {
    body: {
        background: "#aaa",
        borderRadius: 10,
        padding: 8,
        marginTop: 10,
        display: "flex",
        alignItems: "center",
        gap: 6,
    },
    input: {
        margin: 6,
        padding: 6,
        borderRadius: 6,
        border: "1px solid #aaa",
    },
    removeX: {
        marginLeft: "auto",
        background: "none",
        border: "none",
        fontSize: 18,
        cursor: "pointer",
    },

}
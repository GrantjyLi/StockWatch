import React from "react";
import { useState } from "react";
import WatchlistAddAlert from "./WatchlistAddAlert";

export default function AddWatchlistPopup({ addNewWatchlist, onClose }) {
    const [watchlistName, setWatchlistName] = useState([]);
    const [newAlerts, setNewAlerts] = useState([]);

    const addRow = () => {
        setNewAlerts([...newAlerts, { ticker: "", operator: ">=", price: "" }]);
    };

    const removeRow = (index) => {
        setNewAlerts(newAlerts.filter((_, i) => i !== index));
    };

    const updateRow = (index, updatedRow) => {
        setNewAlerts(
        newAlerts.map((row, i) =>
            i === index ? updatedRow : row
        )
        );
    };

    return (
        <div style={styles.overlay}>
            <div style={styles.modal}>
            <div style={styles.header}>Add Watchlist</div>

            <div style={styles.body}>
                <label style={styles.label}>Watchlist Name</label>
                <input style={styles.input} onChange={(e) => setWatchlistName(e.target.value)}/>

                {newAlerts.map((alertRowData, index) => (
                    <WatchlistAddAlert
                        key={index}
                        alertData={alertRowData}
                        onChange={(newRow) => updateRow(index, newRow)}
                        onRemove={() => removeRow(index)}
                    />
                ))}

                <button style={styles.addRow} onClick={addRow}>+</button>
            </div>

            <div style={styles.footer}>
                <button style={styles.footerButton} onClick={()=>addNewWatchlist(watchlistName, newAlerts)}>
                    Add Watchlist
                </button>
                <button style={styles.footerButton} onClick={onClose}>
                    Cancel
                </button>
            </div>

            </div>
        </div>
    );
}

const styles = {
  overlay: {
    position: "fixed",
    inset: 0,
    background: "rgba(0,0,0,0.4)",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    zIndex: 1000,
    fontFamily: "Arial, Helvetica, sans-serif",
  },

  modal: {
    width: 550,
    background: "#ddd",
    borderRadius: 12,
    overflow: "hidden",
  },

  header: {
    background: "#a8f5a2",
    padding: 12,
    textAlign: "center",
    fontWeight: "bold",
  },

  body: {
    padding: 16,
  },

  label: {
    display: "block",
    marginBottom: 6,
  },

  input: {
    margin: 6,
    padding: 6,
    borderRadius: 6,
    border: "1px solid #aaa",
  },

  row: {
    background: "#aaa",
    borderRadius: 10,
    padding: 8,
    marginTop: 10,
    display: "flex",
    alignItems: "center",
    gap: 6,
  },

  addRow: {
    width: "100%",
    marginTop: 12,
    fontSize: 28,
    background: "#aaa",
    border: "none",
    borderRadius: 10,
    cursor: "pointer",
  },

  footer: {
    display: "flex",
    justifyContent: "space-around",
    padding: 16,
  },

  footerButton: {
    padding: "10px 24px",
    borderRadius: 20,
    border: "none",
    cursor: "pointer",
    background: "#fff",
    fontSize: 16,
  },
};

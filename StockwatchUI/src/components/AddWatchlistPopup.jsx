import React from "react";
import { useState } from "react";

export default function AddWatchlistPopup({ onClose }) {
  const [rows, setRows] = useState([{ ticker: "", condition: "", price: "" }]);

  const addRow = () => {
    setRows([...rows, { ticker: "", condition: "", price: "" }]);
  };

  const removeRow = (index) => {
    setRows(rows.filter((_, i) => i !== index));
  };

  return (
    <div style={styles.overlay}>
      <div style={styles.modal}>

        <div style={styles.header}>Add Watchlist</div>

        <div style={styles.body}>
          <label style={styles.label}>Watchlist Name</label>
          <input style={styles.input} />

          <div style={styles.row}>
            <input style={styles.input} placeholder="Ticker" />
            <select style={styles.input} name="cars" id="cars">
                <option value=">=">&gt;=</option>
                <option value="<=">&lt;=</option>
                <option value="=">=</option>
            </select>
            <input style={styles.input} placeholder="Price" type="number" step="any" min="0" name="amount"/>
            <button style={styles.removeX}>✕</button>
          </div>

          <button style={styles.addRow}>+</button>
        </div>

        <div style={styles.footer}>
          <button style={styles.footerButton} onClick={onClose}>
            Remove
          </button>
          <button style={styles.footerButton}>
            Add
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

  removeX: {
    marginLeft: "auto",
    background: "none",
    border: "none",
    fontSize: 18,
    cursor: "pointer",
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

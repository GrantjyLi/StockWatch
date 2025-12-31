import React from "react";

export default function ErrorPopup({ message, onClose }) {
  if (!message) return null; // don't render if no error

  return (
    <div style={styles.overlay}>
      <div style={styles.popup}>
        <div style={styles.header}>Error</div>

        <div style={styles.body}>
          <p>{message}</p>
        </div>

        <button style={styles.button} onClick={onClose}>
          Close
        </button>
      </div>
    </div>
  );
}

const styles = {
  overlay: {
    position: "fixed",
    inset: 0,
    background: "rgba(0,0,0,0.5)",
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    zIndex: 9999,
  },
  popup: {
    background: "#fff",
    padding: "20px 28px",
    borderRadius: 10,
    width: 350,
    textAlign: "center",
    boxShadow: "0 4px 12px rgba(0,0,0,0.2)",
    fontFamily: "Arial, sans-serif",
  },
  header: {
    fontSize: 20,
    fontWeight: "bold",
    marginBottom: 12,
    color: "#b00020",
  },
  body: {
    marginBottom: 20,
  },
  button: {
    padding: "8px 20px",
    borderRadius: 6,
    border: "none",
    background: "#b00020",
    color: "#fff",
    cursor: "pointer",
    fontSize: 14,
  },
};

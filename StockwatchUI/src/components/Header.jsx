import React from "react";

export default function Header({handleLogout}) {
  return (
    <header style={styles.header}>
      <div style={styles.menu}>☰</div>
      <h2>Stockwatch</h2>

      <button style={styles.logoutButton} onClick={handleLogout}>
        Logout
      </button>
    </header>
  );
}

const styles = {
  header: {
    background: "#A6FFA6",
    height: 60,
    display: "flex",
    alignItems: "center",
    padding: "0 20px",
    gap: 20,
    fontFamily: "Arial, Helvetica, sans-serif"
  },
  menu: {
    fontSize: 24,
    cursor: "pointer",
  },
  logoutButton: {
    marginLeft: "auto",
    padding: "6px 14px",
    fontSize: 14,
    fontWeight: "bold",
    cursor: "pointer",
    background: "#fff",
    border: "none",
    borderRadius: 8
  }
};

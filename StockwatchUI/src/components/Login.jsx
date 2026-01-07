import React, { useState } from "react";

export default function Login({ handleLogin }) {
    const [email, setEmail] = useState("");

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!email) return;
        handleLogin(email);
    };

    return (
        <div style={styles.container}>
            <form onSubmit={handleSubmit} style={styles.card}>
                <h2>Sign In</h2>
                <input
                    type="email"
                    placeholder="Email address"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    style={styles.input}
                />
                <button type="submit" style={styles.button}>
                    Continue
                </button>
            </form>
        </div>
    );
}

const styles = {
    container: {
        height: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        background: "#ffffffff",
        fontFamily: "Arial, Helvetica, sans-serif"
    },
    card: {
        background: "#A6FFA6",
        padding: "2rem",
        borderRadius: "8px",
        width: "300px",
        textAlign: "center",
        color: "white",
    },
    input: {
        width: "100%",
        padding: "10px",
        marginBottom: "1rem",
    },
    button: {
        width: "100%",
        padding: "10px",
        cursor: "pointer",
    },
};

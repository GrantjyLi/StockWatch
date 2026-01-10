import React, { useState } from "react";

export default function Login({ handleLogin, handleCreateUser }) {
    const [email, setEmail] = useState("");
    const [isCreating, setIsCreating] = useState(false); // toggle mode

    const handleSubmit = (e) => {
        e.preventDefault();
        if (!email) return;

        if (isCreating) {
            var result = handleCreateUser(email);

            if(result){
                setIsCreating(false)
                setEmail("")
            }
            
        } else {
            handleLogin(email);
        }
    };

    return (
        <div style={styles.container}>
            <form onSubmit={handleSubmit} style={styles.card}>
                <h2 style={styles.title}>
                    {isCreating ? "Create Account" : "Sign In"}
                </h2>

                <input
                    type="email"
                    placeholder="Email address"
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    required
                    style={styles.input}
                />

                <button type="submit" style={styles.button}>
                    {isCreating ? "Create Account" : "Login"}
                </button>

                <button
                    type="button"
                    style={styles.button}
                    onClick={() => setIsCreating(!isCreating)}
                >
                    {isCreating ? "Back to Login" : "Create New User"}
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
    title:{
        color:"black"
    },
    input: {
        width: "100%",
        padding: "10px",
        marginBottom: "1rem",
        boxSizing: "border-box",
    },
    button: {
        width: "100%",
        padding: "10px",
        cursor: "pointer",
        boxSizing: "border-box",
        marginBottom: "0.5rem",
    }
};

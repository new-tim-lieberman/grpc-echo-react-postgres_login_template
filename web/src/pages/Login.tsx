import { useNavigate } from "react-router-dom";
import { useState } from "react";

import api from "../api/client";
import { tokenStore } from "../auth/tokenStore";

export default function Login() {
    const navigate = useNavigate();

    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const login = async () => {
        try {
            const res = await api.post("/api/login", {
                email,
                password,
            });

            const { token, refresh_token } = res.data;

            tokenStore.setTokens(token, refresh_token);

            navigate("/home");
        } catch (e) {
            console.error(e);
        }
    };

    return (
        <div style={{ padding: 40 }}>
            <h1>Login</h1>

            <input
                placeholder="Email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            />

            <br /><br />

            <input
                type="password"
                placeholder="Password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />

            <br /><br />

            <button onClick={login}>
                Login
            </button>
        </div>
    );
}
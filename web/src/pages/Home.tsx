import { useNavigate } from "react-router-dom";

import { tokenStore } from "../auth/tokenStore";

export default function Home() {
    const navigate = useNavigate();

    const logout = () => {
        tokenStore.clear();

        navigate("/");
    };

    return (
        <div style={{ padding: 40 }}>
            <h1>Home Screen</h1>

            <button onClick={logout}>
                Logout
            </button>
        </div>
    );
}
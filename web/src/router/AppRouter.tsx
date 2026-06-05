import {
    BrowserRouter,
    Routes,
    Route,
    Navigate,
} from "react-router-dom";

import Login from "../pages/Login";
import Home from "../pages/Home";

import { tokenStore } from "../auth/tokenStore";

function ProtectedRoute({
                            children,
                        }: {
    children: React.ReactNode;
}) {
    const token = tokenStore.getAccess();

    if (!token) {
        return <Navigate to="/" replace />;
    }

    return children;
}

export default function AppRouter() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Login />} />

                <Route
                    path="/home"
                    element={
                        <ProtectedRoute>
                            <Home />
                        </ProtectedRoute>
                    }
                />
            </Routes>
        </BrowserRouter>
    );
}
import axios from "axios";
import { TokenStore } from "../auth/tokenStore";

const api = axios.create({
    baseURL: "http://localhost:8080",
});

let isRefreshing = false;
let queue: any[] = [];

const processQueue = (error: any, token: string | null = null) => {
    queue.forEach((p) => {
        error ? p.reject(error) : p.resolve(token);
    });
    queue = [];
};

// attach access token
api.interceptors.request.use(async (config) => {
    const token = await TokenStore.getAccess();

    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
});

// refresh logic
api.interceptors.response.use(
    (res) => res,
    async (err) => {
        const original = err.config;

        if (err.response?.status !== 401 || original._retry) {
            return Promise.reject(err);
        }

        original._retry = true;

        if (isRefreshing) {
            return new Promise((resolve, reject) => {
                queue.push({ resolve, reject });
            }).then((token) => {
                original.headers.Authorization = `Bearer ${token}`;
                return api(original);
            });
        }

        isRefreshing = true;

        try {
            const refreshToken = await TokenStore.getRefresh();

            const res = await axios.post(
                "http://localhost:8080/auth/refresh",
                { refreshToken }
            );

            const newAccessToken = res.data.token;

            await TokenStore.setTokens(newAccessToken, refreshToken!);

            processQueue(null, newAccessToken);

            original.headers.Authorization = `Bearer ${newAccessToken}`;
            return api(original);
        } catch (e) {
            processQueue(e, null);
            await TokenStore.clear();
            return Promise.reject(e);
        } finally {
            isRefreshing = false;
        }
    }
);

export default api;
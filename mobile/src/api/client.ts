import axios, {
    AxiosError,
    InternalAxiosRequestConfig,
} from "axios";

import { tokenStore } from "@/auth/tokenStore";

const api = axios.create({
    baseURL: "http://localhost:8080",
});

let isRefreshing = false;

type QueueItem = {
    resolve: (token: string) => void;
    reject: (error: AxiosError | Error) => void;
};

let queue: QueueItem[] = [];

const processQueue = (
    error: AxiosError | Error | null,
    token: string | null = null
) => {
    queue.forEach((p) => {
        if (error) {
            p.reject(error);
        } else if (token) {
            p.resolve(token);
        }
    });

    queue = [];
};

type RetryConfig = InternalAxiosRequestConfig & {
    _retry?: boolean;
};

api.interceptors.request.use(
    async (config: InternalAxiosRequestConfig) => {
        const token = await tokenStore.getAccess();

        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }

        return config;
    }
);

api.interceptors.response.use(
    (response) => response,

    async (error: AxiosError) => {
        const original = error.config as RetryConfig;

        if (
            !error.response ||
            error.response.status !== 401 ||
            original._retry
        ) {
            return Promise.reject(error);
        }

        original._retry = true;

        if (isRefreshing) {
            return new Promise<string>((resolve, reject) => {
                queue.push({ resolve, reject });
            }).then((token) => {
                original.headers.Authorization = `Bearer ${token}`;
                return api(original);
            });
        }

        isRefreshing = true;

        try {
            const refreshToken = await tokenStore.getRefresh();

            const response = await axios.post(
                "http://localhost:8080/api/refresh",
                {
                    refresh_token: refreshToken,
                }
            );

            const newAccessToken = response.data.token;

            await tokenStore.setTokens(
                newAccessToken,
                refreshToken as string
            );

            processQueue(null, newAccessToken);

            original.headers.Authorization = `Bearer ${newAccessToken}`;

            return api(original);
        } catch (err) {
            processQueue(err as AxiosError, null);

            await tokenStore.clear();

            return Promise.reject(err);
        } finally {
            isRefreshing = false;
        }
    }
);

export default api;
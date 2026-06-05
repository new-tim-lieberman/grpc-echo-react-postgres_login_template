import * as SecureStore from "expo-secure-store";

const ACCESS_KEY = "access_token";
const REFRESH_KEY = "refresh_token";

export const tokenStore = {
    async getAccess(): Promise<string | null> {
        return await SecureStore.getItemAsync(ACCESS_KEY);
    },

    async getRefresh(): Promise<string | null> {
        return await SecureStore.getItemAsync(REFRESH_KEY);
    },

    async setTokens(
        access: string,
        refresh: string
    ): Promise<void> {
        await SecureStore.setItemAsync(
            ACCESS_KEY,
            access
        );

        await SecureStore.setItemAsync(
            REFRESH_KEY,
            refresh
        );
    },

    async clear(): Promise<void> {
        await SecureStore.deleteItemAsync(
            ACCESS_KEY
        );

        await SecureStore.deleteItemAsync(
            REFRESH_KEY
        );
    },
};
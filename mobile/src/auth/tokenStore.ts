import * as SecureStore from "expo-secure-store";

const ACCESS = "access_token";
const REFRESH = "refresh_token";

export const TokenStore = {
    getAccess: () => SecureStore.getItemAsync(ACCESS),
    getRefresh: () => SecureStore.getItemAsync(REFRESH),

    setTokens: async (access: string, refresh: string) => {
        await SecureStore.setItemAsync(ACCESS, access);
        await SecureStore.setItemAsync(REFRESH, refresh);
    },

    clear: async () => {
        await SecureStore.deleteItemAsync(ACCESS);
        await SecureStore.deleteItemAsync(REFRESH);
    },
};
import { Stack, useRouter, useSegments } from "expo-router";
import React, { useEffect, useState } from "react";
import { ActivityIndicator, View } from "react-native";

import { tokenStore } from "@/auth/tokenStore";

export default function RootLayout() {
    const router = useRouter();
    const segments = useSegments();

    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const checkAuth = async () => {
            const token = await tokenStore.getAccess();

            //const inAuthGroup = segments[0] === "(auth)";

            if (!token) {
                router.replace("/");
            } else {
                router.replace("/home");
            }

            setLoading(false);
        };

        checkAuth();
    }, []);

    if (loading) {
        return (
            <View
                style={{
                    flex: 1,
                    justifyContent: "center",
                    alignItems: "center",
                }}
            >
                <ActivityIndicator />
            </View>
        );
    }

    return <Stack screenOptions={{ headerShown: false }} />;
}
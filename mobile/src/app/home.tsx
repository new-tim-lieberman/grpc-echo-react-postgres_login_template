import { useRouter } from "expo-router";
import React from "react";
import {
    View,
    Text,
    Button,
} from "react-native";

import { tokenStore } from "@/auth/tokenStore";

export default function HomeScreen() {
    const router = useRouter();

    const logout = async () => {
        await tokenStore.clear();

        router.replace("/");
    };

    return (
        <View
            style={{
                flex: 1,
                justifyContent: "center",
                alignItems: "center",
            }}
        >
            <Text
                style={{
                    fontSize: 28,
                    marginBottom: 20,
                }}
            >
                Home Screen
            </Text>

            <Button title="Logout" onPress={logout} />
        </View>
    );
}
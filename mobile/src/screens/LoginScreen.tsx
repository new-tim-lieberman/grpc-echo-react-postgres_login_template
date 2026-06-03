import React, { useState } from "react";
import { View, TextInput, Button } from "react-native";
import api from "../api/client";
import { TokenStore } from "../auth/tokenStore";

export default function LoginScreen() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    const login = async () => {
        const res = await api.post("/auth/login", {
            email,
            password,
        });

        const { token, refresh_token } = res.data;

        await TokenStore.setTokens(token, refresh_token);

        console.log("logged in");
    };

    return (
        <View style={{ padding: 20 }}>
            <TextInput placeholder="email" value={email} onChangeText={setEmail} />
            <TextInput
                placeholder="password"
                secureTextEntry
                value={password}
                onChangeText={setPassword}
            />
            <Button title="Login" onPress={login} />
        </View>
    );
}
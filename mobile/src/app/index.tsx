import { useRouter } from "expo-router";
import React, { useState } from "react";
import {
  View,
  TextInput,
  Button,
  Text,
} from "react-native";

import api from "../api/client";
import { tokenStore } from "@/auth/tokenStore";

export default function LoginScreen() {
  const router = useRouter();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const login = async () => {
    try {
      const res = await api.post("/api/login", {
        email,
        password,
      });

      const { token, refresh_token } = res.data;

      await tokenStore.setTokens(
          token,
          refresh_token
      );

      router.replace("/home");
    } catch (e) {
      console.log(e);
    }
  };

  return (
      <View
          style={{
            flex: 1,
            justifyContent: "center",
            padding: 20,
          }}
      >
        <Text
            style={{
              fontSize: 28,
              marginBottom: 20,
            }}
        >
          Login
        </Text>

        <TextInput
            placeholder="Email"
            value={email}
            onChangeText={setEmail}
            autoCapitalize="none"
            style={{
              borderWidth: 1,
              marginBottom: 10,
              padding: 10,
            }}
        />

        <TextInput
            placeholder="Password"
            secureTextEntry
            value={password}
            onChangeText={setPassword}
            style={{
              borderWidth: 1,
              marginBottom: 20,
              padding: 10,
            }}
        />

        <Button title="Login" onPress={login} />
      </View>
  );
}
import { router } from "expo-router"
import { Text, View } from "react-native"

import { useSession } from "../features/auth/ctx"

export default function SignInModal() {
  const { signIn } = useSession() || {}
  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
      <Text
        onPress={() => {
          signIn && signIn()
          // Navigate after signing in. You may want to tweak this to ensure sign-in is
          // successful before navigating.
          router.replace("/")
        }}
      >
        register
      </Text>
    </View>
  )
}

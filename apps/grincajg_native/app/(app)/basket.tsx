import { Text } from "@/components"
import { useSession } from "@/features/auth/ctx"
import { NotAuthorizedPageLayout } from "@/layouts/not-authorized"
import { View } from "react-native"

export default function TabBasketScreen() {
  const { session } = useSession() || {}
  if (!session)
    return (
      <NotAuthorizedPageLayout
        pageTitle="Košarica"
        title="Lorem ipsum dolor sit amet"
        descirption="Qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."
      />
    )
  return (
    <View
      style={{
        flex: 1,
        alignItems: "center",
        justifyContent: "center",
      }}
    >
      <Text>kosarica</Text>
    </View>
  )
}

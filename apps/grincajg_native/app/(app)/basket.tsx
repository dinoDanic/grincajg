import { useSession } from "@/features/auth/ctx"
import { NotAuthorizedPageLayout } from "@/layouts/not-authorized"

import { Text, View } from "../../components/Themed"

export default function TabBasketScreen() {
  const { session } = useSession() || {}
  if (!session)
    return (
      <NotAuthorizedPageLayout
        pageTitle="Boro"
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
      <Text style={{}}>kosarica</Text>
    </View>
  )
}

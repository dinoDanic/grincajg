import { Text } from "@/components"
import { useSession } from "@/features/auth/ctx"
import { NotAuthorizedPageLayout } from "@/layouts/not-authorized"
import { StyleSheet, View } from "react-native"

export default function TabInboxScreen() {
  const { session } = useSession() || {}
  if (!session)
    return (
      <NotAuthorizedPageLayout
        pageTitle="Poruke"
        title="Lorem ipsum dolor sit amet"
        descirption="Qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."
      />
    )
  return (
    <View style={styles.container}>
      <Text>inbox</Text>
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
  },
  title: {
    fontSize: 20,
    fontWeight: "bold",
  },
  separator: {
    marginVertical: 30,
    height: 1,
    width: "80%",
  },
})

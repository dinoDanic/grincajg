import { NotAuthorizedPageLayout } from "@/layouts/not-authorized"
import { StyleSheet } from "react-native"

import { Text, View } from "../../components/Themed"
import { useSession } from "../../features/auth/ctx"

export default function TabProfileScreen() {
  const { session } = useSession() || {}
  if (!session)
    return (
      <NotAuthorizedPageLayout
        pageTitle="Porfil"
        title="Lorem ipsum dolor sit amet"
        descirption="Qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."
      />
    )
  return (
    <View style={styles.container}>
      <Text>profile</Text>
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

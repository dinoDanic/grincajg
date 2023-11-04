import { Colors } from "@/constants"
import { Tabs } from "expo-router"
import { MessageCircleIcon, SearchIcon, ShoppingBasketIcon, UserCircleIcon } from "lucide-react-native"

import { Text } from "../../components/Themed"
import { useSession } from "../../features/auth/ctx"

export default function TabLayout() {
  const { isLoading } = useSession() || {}

  if (isLoading) {
    return <Text>Loading..</Text>
  }

  // if (!session) {
  //   return <Redirect href="/sign-in" />
  // }

  return (
    <Tabs
      screenOptions={{
        tabBarActiveTintColor: Colors.primary,
        headerShown: false,
      }}
    >
      <Tabs.Screen
        name="basket"
        options={{
          title: "Kosarica",
          tabBarIcon: ({ color }) => <ShoppingBasketIcon color={color} />,
        }}
      />
      <Tabs.Screen
        name="index"
        options={{
          title: "Pretrazi",
          tabBarIcon: ({ color }) => <SearchIcon color={color} />,
        }}
      />
      <Tabs.Screen
        name="inbox"
        options={{
          title: "Poruke",
          tabBarIcon: ({ color }) => <MessageCircleIcon color={color} />,
        }}
      />
      <Tabs.Screen
        name="profile"
        options={{
          title: "Profil",
          tabBarIcon: ({ color }) => <UserCircleIcon color={color} />,
        }}
      />
    </Tabs>
  )
}

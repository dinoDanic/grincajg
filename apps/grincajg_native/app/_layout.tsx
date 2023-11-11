import { useEffect } from "react"
import { SessionProvider } from "@/features/auth/ctx"
import { QueryProvider } from "@/lib"
import FontAwesome from "@expo/vector-icons/FontAwesome"
import { DarkTheme, DefaultTheme, ThemeProvider } from "@react-navigation/native"
import { useFonts } from "expo-font"
import { SplashScreen, Stack } from "expo-router"
import { useColorScheme } from "react-native"

export {
  // Catch any errors thrown by the Layout component.
  ErrorBoundary,
} from "expo-router"

export const unstable_settings = {
  // Ensure that reloading on `/modal` keeps a back button present.
  initialRouteName: "(app)",
}

// Prevent the splash screen from auto-hiding before asset loading is complete.
SplashScreen.preventAutoHideAsync()

export default function RootLayout() {
  const [loaded, error] = useFonts({
    SpaceMono: require("../assets/fonts/SpaceMono-Regular.ttf"),
    ...FontAwesome.font,
  })

  // Expo Router uses Error Boundaries to catch errors in the navigation tree.
  useEffect(() => {
    if (error) throw error
  }, [error])

  useEffect(() => {
    if (loaded) {
      SplashScreen.hideAsync()
    }
  }, [loaded])

  if (!loaded) {
    return null
  }

  return <RootLayoutNav />
}

function RootLayoutNav() {
  const colorScheme = useColorScheme()

  return (
    <ThemeProvider value={colorScheme === "dark" ? DarkTheme : DefaultTheme}>
      <QueryProvider>
        <SessionProvider>
          <Stack screenOptions={{ headerShown: false }}>
            {/*   <Stack.Screen */}
            {/*     name="(app)" */}
            {/*     options={{ headerShown: false, presentation: "modal" }} */}
            {/*   /> */}
            <Stack.Screen name="sign-in-modal" options={{ presentation: "modal" }} />
            <Stack.Screen name="register-modal" options={{ presentation: "modal" }} />
          </Stack>
        </SessionProvider>
      </QueryProvider>
    </ThemeProvider>
  )
}

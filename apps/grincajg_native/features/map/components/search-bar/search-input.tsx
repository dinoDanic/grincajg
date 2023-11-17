import React from "react"
import { Text } from "@/components"
import { Colors, Spacing } from "@/constants"
import { SearchIcon, SlidersHorizontal } from "lucide-react-native"
import { StyleSheet, View } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"

export const SearchInput = () => {
  return (
    <View style={{ ...styles.container, backgroundColor: "white" }}>
      <View style={styles.inputwrap}>
        <SearchIcon size={20} color={Colors.dark} strokeWidth={3} />
        <View>
          <Text size="sm" weight="bold">
            Pretrazi
          </Text>
          <Text size="sm">Lokacija, mjesto, naselje</Text>
        </View>
      </View>
      <TouchableOpacity style={{ justifyContent: "center", alignItems: "center", flex: 1, padding: Spacing.sm }}>
        <SlidersHorizontal color={Colors.dark} size={20} style={{ marginTop: 3 }} />
      </TouchableOpacity>
    </View>
  )
}

const styles = StyleSheet.create({
  wrapper: {
    position: "absolute",
    width: "100%",
  },
  container: {
    width: "100%",
    alignSelf: "center",
    padding: Spacing.md,
    flexDirection: "row",
    gap: Spacing.sm,
  },
  inputwrap: {
    borderWidth: 1,
    borderColor: Colors.input,
    padding: Spacing.md,
    shadowColor: "#000",
    backgroundColor: "white",
    borderRadius: 999,
    shadowOpacity: 0.1,
    shadowRadius: 5,
    flexDirection: "row",
    gap: Spacing.md,
    flexGrow: 1,
    alignItems: "center",
  },
})

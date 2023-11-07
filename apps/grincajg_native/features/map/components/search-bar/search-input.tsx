import React, { useState } from "react"
import { Colors, Spacing } from "@/constants"
import { SearchIcon } from "lucide-react-native"
import { StyleSheet, TextInput, View } from "react-native"

export const SearchInput = () => {
  const [text, onChangeText] = useState("")
  return (
    <View style={{ ...styles.container, backgroundColor: "white" }}>
      <View style={styles.inputwrap}>
        <SearchIcon color={Colors.primary} />
        <TextInput
          onChangeText={(newText) => onChangeText(newText)}
          value={text}
          placeholder="Pretrazi"
          style={{
            color: Colors.primary,
          }}
        />
      </View>
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
  },
})

import { Border, Spacing } from "@/constants"
import { StyleSheet, TextInput, TextInputProps } from "react-native"

export const Input = (props: TextInputProps) => {
  return <TextInput style={styles.input} {...props} />
}

const styles = StyleSheet.create({
  input: {
    borderColor: "gray",
    borderWidth: 1,
    width: "100%",
    padding: Spacing.sm,
    borderRadius: Border.md,
  },
})

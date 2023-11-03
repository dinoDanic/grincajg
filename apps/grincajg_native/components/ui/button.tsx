import React, { PropsWithChildren } from "react"
import { Border, Spacing } from "@/constants"
import { TouchableOpacity } from "react-native-gesture-handler"

import { Text } from "../Themed"

export const Button = ({ children }: PropsWithChildren) => {
  return (
    <TouchableOpacity style={{ backgroundColor: "green", padding: Spacing.md, borderRadius: Border.md, }}>
      <Text>{children}</Text>
    </TouchableOpacity>
  )
}

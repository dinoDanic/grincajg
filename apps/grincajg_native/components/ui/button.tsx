import React, { PropsWithChildren } from "react"
import { Border, Colors, Spacing } from "@/constants"
import { StyleProp, TextStyle, ViewStyle } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"

import { Text } from "../Themed"

type Variants = "primary" | "link"

type Props = PropsWithChildren & {
  style?: StyleProp<ViewStyle>
  variants?: Variants
}

export const Button = ({ children, variants = "primary", ...props }: Props) => {
  const containerStyleVariants: Record<Variants, StyleProp<ViewStyle>> = {
    primary: {
      backgroundColor: Colors.primary,
    },
    link: {},
  }
  const textStyleVariants: Record<Variants, StyleProp<TextStyle>> = {
    primary: {},
    link: {
      color: Colors.primary,
      fontWeight: "700",
    },
  }
  return (
    <TouchableOpacity
      style={[
        {
          padding: Spacing.md,
          borderRadius: Border.md,
        },
        containerStyleVariants[variants],
        props.style,
      ]}
    >
      <Text style={[{ color: "white", alignSelf: "center" }, textStyleVariants[variants]]}>{children}</Text>
    </TouchableOpacity>
  )
}

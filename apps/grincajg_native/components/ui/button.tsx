import React, { PropsWithChildren } from "react"
import { Border, Colors, Spacing } from "@/constants"
import { StyleProp, TextStyle, ViewStyle } from "react-native"
import { TouchableOpacity } from "react-native-gesture-handler"

import { Text } from "../Themed"

type Variants = "primary" | "link" | "secondary" | "ghost"

type Sizes = "icon" | "default"

type Props = PropsWithChildren & {
  style?: StyleProp<ViewStyle>
  variants?: Variants
  size?: Sizes
}

export const Button = ({ children, variants = "primary", size = "default", ...props }: Props) => {
  const containerStyleVariants: Record<Variants, StyleProp<ViewStyle>> = {
    primary: {
      backgroundColor: Colors.primary,
    },
    secondary: {
      backgroundColor: "white",
      borderWidth: 1,
    },
    link: {},
    ghost: {},
  }
  const textStyleVariants: Record<Variants, StyleProp<TextStyle>> = {
    primary: {},
    link: {
      color: Colors.primary,
      fontWeight: "700",
    },
    secondary: {
      color: Colors.primary,
      fontWeight: "700",
    },
    ghost: {
      color: Colors.primary,
      fontWeight: "700",
    },
  }
  const sizesVariants: Record<Sizes, StyleProp<ViewStyle>> = {
    icon: {
      width: 40,
      height: 40,
      justifyContent: "center",
      alignItems: "center",
      padding: 0,
    },
    default: {},
  }
  return (
    <TouchableOpacity
      style={[
        {
          padding: Spacing.md,
          borderRadius: Border.md,
        },
        containerStyleVariants[variants],
        sizesVariants[size],
        props.style,
      ]}
    >
      <Text style={[{ color: "white", alignSelf: "center" }, textStyleVariants[variants]]}>{children}</Text>
    </TouchableOpacity>
  )
}

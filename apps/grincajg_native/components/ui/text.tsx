import React from "react"
import { Colors } from "@/constants"
import { Text as TextNative, TextStyle, type TextProps } from "react-native"

type Sizes = "default" | "sm" | "xs"
export type TextWeight = "default" | "bold" | "bolder" | "oh tha boldesta"
export type TextTone = "primary" | "default" | "muted"

type Props = TextProps & {
  size?: Sizes
  weight?: TextWeight
  tone?: TextTone
}

const sizesVariants: Record<Sizes, TextStyle["fontSize"]> = {
  default: 14,
  sm: 10,
  xs: 8,
}

const weightVariants: Record<TextWeight, TextStyle["fontWeight"]> = {
  default: "300",
  bold: "bold",
  bolder: "800",
  "oh tha boldesta": "900",
}

const colorVariants: Record<TextTone, TextStyle["color"]> = {
  default: Colors.dark,
  primary: Colors.primary,
  muted: Colors["primary-foreground"],
}

export const Text = ({ size = "default", weight = "default", tone = "default", ...props }: Props) => {
  return (
    <TextNative
      style={{ fontSize: sizesVariants[size], fontWeight: weightVariants[weight], color: colorVariants[tone] }}
      {...props}
    />
  )
}

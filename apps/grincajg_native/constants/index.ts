export const Spacing = {
  xxs: 2,
  xs: 4,
  sm: 8,
  md: 12,
  lg: 16,
  xl: 20,
  "2xl": 24,
  "3xl": 32,
  "5xl": 40,
}

export const Border = {
  md: 6,
  xl: 12,
  lg: 24,
}

const Pallete = {
  creamyWhite: {
    foreground: "hsl(41, 100%, 97%)",
    background: "hsl(41, 100%, 96%)",
  },
  honeyGold: {
    foreground: "hsl(36, 100%, 65%)",
    background: "hsl(36, 100%, 64%)",
  },
  cherryTomatoRed: {
    foreground: "hsl(2, 87%, 63%)",
    background: "hsl(2, 87%, 62%)",
  },
  sageGreen: {
    foreground: "hsl(82, 26%, 65%)",
    background: "hsl(82, 26%, 64%)",
  },
  butteryYellow: {
    foreground: "hsl(47, 100%, 70%)",
    background: "hsl(47, 100%, 69%)",
  },
  toastedAlmond: {
    foreground: "hsl(34, 30%, 69%)",
    background: "hsl(34, 30%, 68%)",
  },
  caramelDrizzle: {
    foreground: "hsla(21, 34%, 45%, 0.8)",
    background: "hsl(21, 34%, 45%)",
  },
}

export const Colors = {
  primary: Pallete.caramelDrizzle.background,
  input: "#eee",
  "primary-foreground": Pallete.caramelDrizzle.foreground,
  accents: Pallete.sageGreen.foreground,
  gray: "#999",
  dark: "#333",
  muted: "#f1f5f9",
  mutedForeground: "#64748b",
}

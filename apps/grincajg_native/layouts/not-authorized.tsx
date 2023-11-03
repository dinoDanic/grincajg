import React from "react"
import { Spacing } from "@/constants"
import { SafeAreaView, StyleSheet } from "react-native"

import { Button } from "@/components/ui/button"
import { Text, View } from "@/components/Themed"

type Props = {
  pageTitle: string
  title: string
  descirption: string
}

export const NotAuthorizedPageLayout = ({ pageTitle, title, descirption }: Props) => {
  return (
    <>
      <SafeAreaView style={{ backgroundColor: "white" }} />
      <View style={styles.container}>
        <Text style={styles.pageTitle}>{pageTitle}</Text>
        <View style={styles.subCont}>
          <Text style={styles.title}>{title}</Text>
          <Text style={styles.descirption}>{descirption}</Text>
        </View>
        <View style={styles.subCont}>
          <Button>Login</Button>
        </View>
      </View>
    </>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: Spacing.lg,
    gap: Spacing.xl,
  },
  pageTitle: {
    fontSize: 25,
    fontWeight: "600",
  },
  title: {
    fontSize: 15,
    fontWeight: "600",
  },
  descirption: {
    fontSize: 15,
  },
  subCont: {
    padding: Spacing.md,
    gap: Spacing.sm,
  },
})

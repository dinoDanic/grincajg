import React from "react"
import { Button, Text } from "@/components"
import { Spacing } from "@/constants"
import { Link } from "expo-router"
import { SafeAreaView, StyleSheet, View } from "react-native"

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
        <View style={styles.actions}>
          <Link href="/sign-in-modal">
            <Button style={{ width: 100 }}>Prijava</Button>
          </Link>
          <Link href="/register-modal">
            <Button variants="link" style={{ width: 120 }}>
              Registracija
            </Button>
          </Link>
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
  actions: {
    padding: Spacing.md,
    gap: Spacing.sm,
    flexDirection: "row",
  },
})

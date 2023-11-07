import React from "react"
import { Colors, Spacing } from "@/constants"
import { Category } from "@/generated/graphql"
import { Text, View } from "react-native"

import { Button } from "@/components/ui/button"

import { iconsByCategoryId } from "../helpers/icons-by-category"

type Props = {
  category: Category
}

export const CategoryIconCard = (props: Props) => {
  return (
    <Button variants="ghost">
      <View style={{ gap: Spacing.sm, alignItems: "center" }}>
        {iconsByCategoryId(Number(props.category.id))}
        <Text style={{ fontSize: 10, color: Colors.gray, fontWeight: "500" }}> {props.category.name}</Text>
      </View>
    </Button>
  )
  // return (
  //   <View style={styles.crad}>
  //     <Text>{props.category.name}</Text>
  //   </View>
  // )
}

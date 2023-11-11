import React from "react"
import { Text } from "@/components"
import { Colors, Spacing } from "@/constants"
import { useMapContext } from "@/features/map"
import { Category } from "@/generated/graphql"
import { View } from "react-native"

import { Button } from "@/components/ui/button"
import { TextTone, TextWeight } from "@/components/ui/text"

import { iconsByCategoryId } from "../helpers/icons-by-category"

type Props = {
  category: Category
}

export const CategoryIconCard = (props: Props) => {
  const setSelectedCategory = useMapContext((store) => store.setSelectedCategory)
  const selectedCategory = useMapContext((store) => store.selectedCategory)

  const isActive = selectedCategory === props.category.id

  const color = isActive ? Colors.primary : Colors.gray

  const nameWeight: TextWeight = isActive ? "bold" : "default"

  const nameTone: TextTone = isActive ? "primary" : "default"

  return (
    <Button onPress={() => setSelectedCategory(props.category.id)} variants="ghost">
      <View style={{ gap: Spacing.sm, alignItems: "center", minWidth: 30 }}>
        {iconsByCategoryId(Number(props.category.id), color)}
        <Text weight={nameWeight} tone={nameTone}>
          {props.category.name}
        </Text>
        {/* {isActive && ( */}
        {/*   <View */}
        {/*     style={{ */}
        {/*       borderBottomWidth: 1, */}
        {/*       width: "100%", */}
        {/*       borderColor: Colors.primary, */}
        {/*       position: "absolute", */}
        {/*       bottom: 0, */}
        {/*     }} */}
        {/*   /> */}
        {/* )} */}
      </View>
    </Button>
  )
}

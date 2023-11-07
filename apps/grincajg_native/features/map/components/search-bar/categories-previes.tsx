import React from "react"
import { CategoryIconCard } from "@/features/categories/components/category-icon-card"
import { GetMainCategoriesDocument } from "@/generated/graphql"
import { _client } from "@/lib"
import { useQuery } from "@tanstack/react-query"
import { FlatList } from "react-native"

export const CategoriesPrevies = () => {
  const { data } = useQuery({
    queryKey: ["GetMainCategoriesDocument"],
    queryFn: () => _client.request(GetMainCategoriesDocument),
  })

  if (data?.categories) {
    return (
      <FlatList
        horizontal
        data={data.categories}
        renderItem={(item) => item.item && <CategoryIconCard category={item.item} />}
      />
    )
  }
}

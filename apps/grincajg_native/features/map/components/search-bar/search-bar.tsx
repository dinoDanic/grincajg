import { SafeAreaView, StyleSheet, View } from "react-native"

import { CategoriesPrevies } from "./categories-previes"
import { SearchInput } from "./search-input"

export const SearchBar = () => {
  return (
    <>
      <View style={styles.wrapper}>
        <SafeAreaView />
        <SearchInput />
        <CategoriesPrevies />
      </View>
    </>
  )
}

const styles = StyleSheet.create({
  wrapper: {
    position: "absolute",
    width: "100%",
    backgroundColor: "white",
  },
})

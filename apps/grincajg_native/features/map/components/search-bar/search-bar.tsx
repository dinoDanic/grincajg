import { useState } from "react";
import { SafeAreaView, StyleSheet, View } from "react-native";
import { TextInput } from "react-native-gesture-handler";

import { useThemeColor } from "../../../../components/Themed";
import Colors from "../../../../constants/Colors";

export const SearchBar = () => {
  const [text, onChangeText] = useState("");

  const background = useThemeColor(
    { light: Colors.light.background, dark: Colors.dark.background },
    "background",
  );
  const inputColor = useThemeColor(
    { light: Colors.light.text, dark: Colors.dark.text },
    "text",
  );
  return (
    <>
      <View style={styles.wrapper}>
        <SafeAreaView />
        <View style={{ ...styles.container, backgroundColor: background }}>
          <TextInput
            onChangeText={(newText) => onChangeText(newText)}
            value={text}
            placeholder="Pretrazi"
            style={{ color: inputColor }}
          />
        </View>
      </View>
    </>
  );
};

const styles = StyleSheet.create({
  wrapper: {
    position: "absolute",
    width: "100%",
  },
  container: {
    borderRadius: 15,
    padding: 20,
    width: "80%",
    alignSelf: "center",
  },
});

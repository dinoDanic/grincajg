import React from "react";
import { StyleSheet, View } from "react-native";
import MapView, { Marker } from "react-native-maps";
import { SearchBar } from "../../features/map/components/search-bar/search-bar";

export default function App() {
  return (
    <View style={styles.container}>
      <MapView
        // onPress={(e) => console.log(e)}
        style={styles.map}
        // onRegionChange={(e) => console.log(e)}
      >
        <Marker
          draggable
          coordinate={{
            latitude: 45.62839378272178,
            longitude: 15.951489423726455,
          }}
          title="title"
          description="descirption"
        />
      </MapView>

      <SearchBar />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  map: {
    width: "100%",
    height: "100%",
  },
});

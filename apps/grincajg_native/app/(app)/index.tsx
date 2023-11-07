import React from "react"
import { Colors } from "@/constants"
import { useMapView } from "@/features/map/hooks/useMapView"
import { StyleSheet, View } from "react-native"
import MapView, { Marker } from "react-native-maps"

import { Text } from "@/components/Themed"

import { SearchBar } from "../../features/map/components/search-bar/search-bar"

export default function App() {
  const { query, onRegionChangeComplete } = useMapView()
  const { data, isFetching } = query

  return (
    <View style={styles.container}>
      <MapView
        camera={{ pitch: 0, center: { latitude: 45.7, longitude: 16 }, zoom: 9, heading: 0 }}
        minZoomLevel={9}
        style={styles.map}
        provider="google"
        onRegionChangeComplete={onRegionChangeComplete}
      >
        {data?.getOrganizationsOnMap?.map((org) => {
          if (!org?.latitude || !org.longitude) return null
          return (
            <Marker
              key={org.id}
              coordinate={{ longitude: org?.longitude, latitude: org?.latitude }}
              title={`organization id: ${org.id}`}
              pinColor={Colors.primary}
              description="vazin"
            />
          )
        })}
      </MapView>
      {isFetching && (
        <View
          style={{ position: "absolute", top: 120, left: 20, backgroundColor: "white", padding: 10, borderRadius: 10 }}
        >
          <Text>LOADING</Text>
        </View>
      )}
      <SearchBar />
    </View>
  )
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  map: {
    width: "100%",
    height: "100%",
  },
})

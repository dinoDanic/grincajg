import { useState } from "react"
import { GetOrganizationsOnMapDocument, GetOrganizationsOnMapInput } from "@/generated/graphql"
import { _client } from "@/lib"
import { keepPreviousData, useQuery } from "@tanstack/react-query"
import { Region } from "react-native-maps"

export const useMapView = () => {
  const [input, setInput] = useState<GetOrganizationsOnMapInput>({
    southwest: { latitude: 0, longitude: 0 },
    northeast: { longitude: 0, latitude: 0 },
  })

  const query = useQuery({
    queryKey: ["map", input],
    queryFn: async () => _client.request(GetOrganizationsOnMapDocument, { input }),
    placeholderData: keepPreviousData,
  })

  const onRegionChangeComplete = (region: Region) => {
    const { latitude, longitude, latitudeDelta, longitudeDelta } = region
    const northeast = {
      latitude: latitude + latitudeDelta / 2,
      longitude: longitude + longitudeDelta / 2,
    }
    const southwest = {
      latitude: latitude - latitudeDelta / 2,
      longitude: longitude - longitudeDelta / 2,
    }
    setInput({ northeast, southwest })
  }

  return { query, onRegionChangeComplete }
}

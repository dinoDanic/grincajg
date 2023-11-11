import { create } from "zustand"

interface MapContext {
  selectedCategory: string | null
  setSelectedCategory: (category: string) => void
}

export const useMapContext = create<MapContext>()((set) => ({
  selectedCategory: null,
  setSelectedCategory: (cat) => set({ selectedCategory: cat }),
}))

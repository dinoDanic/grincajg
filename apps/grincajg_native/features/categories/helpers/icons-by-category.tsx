import { Colors } from "@/constants"
import {
  AppleIcon,
  BananaIcon,
  BeanIcon,
  CarrotIcon,
  CoffeeIcon,
  CroissantIcon,
  EggIcon,
  GlassWaterIcon,
  GrapeIcon,
  HopIcon,
  LeafyGreenIcon,
  MilkIcon,
  SoupIcon,
  TreePineIcon,
} from "lucide-react-native"

const color = Colors.gray

export const iconsByCategoryId = (id: number) => {
  switch (id) {
    case 1:
      return <EggIcon color={color} />
    case 2:
      return <TreePineIcon color={color} />
    case 3:
      return <MilkIcon color={color} />
    case 4:
      return <CroissantIcon color={color} />
    case 5:
      return <GlassWaterIcon color={color} />
    case 6:
      return <CoffeeIcon color={color} />
    case 7:
      return <CarrotIcon color={color} />
    case 8:
      return <BananaIcon color={color} />
    case 9:
      return <BeanIcon color={color} />
    case 10:
      return <GrapeIcon color={color} />
    case 11:
      return <HopIcon color={color} />
    case 12:
      return <LeafyGreenIcon color={color} />
    case 13:
      return <SoupIcon color={color} />
    default:
      return <MilkIcon color={color} />
  }
}

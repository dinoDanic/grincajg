import {
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
  LucideProps,
  MilkIcon,
  SoupIcon,
  TreePineIcon,
} from "lucide-react-native"

export const iconsByCategoryId = (id: number, color: string) => {
  const iconsProps: LucideProps = {
    color,
  }
  switch (id) {
    case 1:
      return <EggIcon {...iconsProps} />
    case 2:
      return <TreePineIcon {...iconsProps} />
    case 3:
      return <MilkIcon {...iconsProps} />
    case 4:
      return <CroissantIcon {...iconsProps} />
    case 5:
      return <GlassWaterIcon {...iconsProps} />
    case 6:
      return <CoffeeIcon {...iconsProps} />
    case 7:
      return <CarrotIcon {...iconsProps} />
    case 8:
      return <BananaIcon {...iconsProps} />
    case 9:
      return <BeanIcon {...iconsProps} />
    case 10:
      return <GrapeIcon {...iconsProps} />
    case 11:
      return <HopIcon {...iconsProps} />
    case 12:
      return <LeafyGreenIcon {...iconsProps} />
    case 13:
      return <SoupIcon {...iconsProps} />
    default:
      return <MilkIcon {...iconsProps} />
  }
}

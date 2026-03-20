package combat

var ArmorEfficiency = map[AmmoType]map[ArmorType]float64{
	AmmoNormal: {ArmorLight: 1.0, ArmorMedium: 0.8, ArmorHeavy: 0.6},
	AmmoHE:     {ArmorLight: 1.2, ArmorMedium: 1.1, ArmorHeavy: 0.9},
	AmmoAP:     {ArmorLight: 0.8, ArmorMedium: 1.2, ArmorHeavy: 1.3},
}

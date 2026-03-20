package combat

import (
	"math"
)

func CalculateDamage(input CombatInput) (CombatResult, error) {
	if err := ValidateInput(input); err != nil {
		return CombatResult{}, err
	}

	armorMod := ArmorEfficiency[input.AmmoType][input.ArmorType]

	defenseRatio := 1.0 - float64(input.Defense)/1000.0
	defenseRatio = math.Max(defenseRatio, MinDamageReductionFloor)

	baseDamage := float64(input.Attack)
	finalDamage := math.Floor(baseDamage * armorMod * defenseRatio)

	return CombatResult{
		BaseDamage:     baseDamage,
		DefenseReduced: defenseRatio,
		FinalDamage:    int(finalDamage),
	}, nil
}

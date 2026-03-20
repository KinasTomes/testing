package main

import (
	"fmt"
	"log"

	"combat-game/internal/combat"
)

func main() {
	input := combat.CombatInput{
		Attack:    320,
		Defense:   220,
		AmmoType:  combat.AmmoAP,
		ArmorType: combat.ArmorHeavy,
	}

	result, err := combat.CalculateDamage(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Damage: %d (base=%.2f, defenseRatio=%.2f)\n", result.FinalDamage, result.BaseDamage, result.DefenseReduced)
}

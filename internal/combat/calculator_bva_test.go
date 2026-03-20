package combat

import (
	"testing"
)

// TestCalculateDamage_AttackBoundary kiểm thử giá trị biên của tham số Attack.
// Các tham số còn lại cố định: Defense=500, AmmoType=Normal, ArmorType=Light
// armorMod = 1.0, defenseRatio = 1 - 500/1000 = 0.5
func TestCalculateDamage_AttackBoundary(t *testing.T) {
	tests := []struct {
		id          string
		name        string
		input       CombatInput
		wantFinal   int
		wantBase    float64
		wantDefRatio float64
		wantErr     bool
	}{
		{
			id:   "TC01",
			name: "Attack = Min (1)",
			input: CombatInput{Attack: 1, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// floor(1 * 1.0 * 0.5) = floor(0.5) = 0
			wantBase:    1,
			wantDefRatio: 0.5,
			wantFinal:   0,
		},
		{
			id:   "TC02",
			name: "Attack = Min+ (2)",
			input: CombatInput{Attack: 2, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// floor(2 * 1.0 * 0.5) = floor(1.0) = 1
			wantBase:    2,
			wantDefRatio: 0.5,
			wantFinal:   1,
		},
		{
			id:   "TC03",
			name: "Attack = Norm (2500)",
			input: CombatInput{Attack: 2500, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// floor(2500 * 1.0 * 0.5) = 1250
			wantBase:    2500,
			wantDefRatio: 0.5,
			wantFinal:   1250,
		},
		{
			id:   "TC04",
			name: "Attack = Max- (4999)",
			input: CombatInput{Attack: 4999, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// floor(4999 * 1.0 * 0.5) = floor(2499.5) = 2499
			wantBase:    4999,
			wantDefRatio: 0.5,
			wantFinal:   2499,
		},
		{
			id:   "TC05",
			name: "Attack = Max (5000)",
			input: CombatInput{Attack: 5000, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// floor(5000 * 1.0 * 0.5) = 2500
			wantBase:    5000,
			wantDefRatio: 0.5,
			wantFinal:   2500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id+"_"+tt.name, func(t *testing.T) {
			got, err := CalculateDamage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("[%s] wantErr=%v, got err=%v", tt.id, tt.wantErr, err)
			}
			if err != nil {
				return
			}
			if got.FinalDamage != tt.wantFinal {
				t.Errorf("[%s] FinalDamage: want %d, got %d", tt.id, tt.wantFinal, got.FinalDamage)
			}
			if got.BaseDamage != tt.wantBase {
				t.Errorf("[%s] BaseDamage: want %.0f, got %.0f", tt.id, tt.wantBase, got.BaseDamage)
			}
			if got.DefenseReduced != tt.wantDefRatio {
				t.Errorf("[%s] DefenseReduced: want %.4f, got %.4f", tt.id, tt.wantDefRatio, got.DefenseReduced)
			}
		})
	}
}

// TestCalculateDamage_DefenseBoundary kiểm thử giá trị biên của tham số Defense.
// Các tham số còn lại cố định: Attack=1000, AmmoType=Normal, ArmorType=Light
// armorMod = 1.0; defenseRatio = max(1 - Defense/1000, 0.1)
func TestCalculateDamage_DefenseBoundary(t *testing.T) {
	tests := []struct {
		id           string
		name         string
		input        CombatInput
		wantFinal    int
		wantBase     float64
		wantDefRatio float64
		wantErr      bool
	}{
		{
			id:   "TC06",
			name: "Defense = Min (0)",
			input: CombatInput{Attack: 1000, Defense: 0, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = max(1 - 0/1000, 0.1) = 1.0
			// floor(1000 * 1.0 * 1.0) = 1000
			wantBase:     1000,
			wantDefRatio: 1.0,
			wantFinal:    1000,
		},
		{
			id:   "TC07",
			name: "Defense = Min+ (1)",
			input: CombatInput{Attack: 1000, Defense: 1, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = 1 - 1/1000 = 0.999
			// floor(1000 * 1.0 * 0.999) = 999
			wantBase:     1000,
			wantDefRatio: 0.999,
			wantFinal:    999,
		},
		{
			id:   "TC08",
			name: "Defense = Norm (500)",
			input: CombatInput{Attack: 1000, Defense: 500, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = 1 - 500/1000 = 0.5
			// floor(1000 * 1.0 * 0.5) = 500
			wantBase:     1000,
			wantDefRatio: 0.5,
			wantFinal:    500,
		},
		{
			id:   "TC09",
			name: "Defense = Max- (999)",
			input: CombatInput{Attack: 1000, Defense: 999, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = max(1 - 999/1000, 0.1) = max(0.001, 0.1) = 0.1
			// floor(1000 * 1.0 * 0.1) = 100
			wantBase:     1000,
			wantDefRatio: 0.1,
			wantFinal:    100,
		},
		{
			id:   "TC10",
			name: "Defense = Max (1000)",
			input: CombatInput{Attack: 1000, Defense: 1000, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = max(1 - 1000/1000, 0.1) = max(0.0, 0.1) = 0.1
			// floor(1000 * 1.0 * 0.1) = 100
			wantBase:     1000,
			wantDefRatio: 0.1,
			wantFinal:    100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id+"_"+tt.name, func(t *testing.T) {
			got, err := CalculateDamage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("[%s] wantErr=%v, got err=%v", tt.id, tt.wantErr, err)
			}
			if err != nil {
				return
			}
			if got.FinalDamage != tt.wantFinal {
				t.Errorf("[%s] FinalDamage: want %d, got %d", tt.id, tt.wantFinal, got.FinalDamage)
			}
			if got.BaseDamage != tt.wantBase {
				t.Errorf("[%s] BaseDamage: want %.0f, got %.0f", tt.id, tt.wantBase, got.BaseDamage)
			}
			if got.DefenseReduced != tt.wantDefRatio {
				t.Errorf("[%s] DefenseReduced: want %.4f, got %.4f", tt.id, tt.wantDefRatio, got.DefenseReduced)
			}
		})
	}
}

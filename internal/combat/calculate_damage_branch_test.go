package combat

import (
	"testing"
)

func TestCalculateDamage_BranchCoverage(t *testing.T) {
	tests := []struct {
		id      string
		path    string
		input   CombatInput
		want    CombatResult
		wantErr bool
	}{
		{
			id:   "CD1",
			path: "1, 2T, 7",
			input: CombatInput{
				Attack:    -1,
				Defense:   -1,
				AmmoType:  AmmoNormal,
				ArmorType: ArmorLight,
			},
			want:    CombatResult{BaseDamage: 0, DefenseReduced: 0, FinalDamage: 0},
			wantErr: true,
		},
		{
			id:   "CD2",
			path: "1, 2F, 3, 4F, 5, 6",
			input: CombatInput{
				Attack:    1000,
				Defense:   200,
				AmmoType:  AmmoHE,
				ArmorType: ArmorLight,
			},
			want:    CombatResult{BaseDamage: 1000, DefenseReduced: 0.8, FinalDamage: 960},
			wantErr: false,
		},
		{
			id:   "CD3",
			path: "1, 2F, 3, 4T, 8, 5, 6",
			input: CombatInput{
				Attack:    360,
				Defense:   960,
				AmmoType:  AmmoAP,
				ArmorType: ArmorHeavy,
			},
			want:    CombatResult{BaseDamage: 360, DefenseReduced: 0.1, FinalDamage: 46},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, err := CalculateDamage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("[%s] CalculateDamage() error = %v, wantErr %v", tt.id, err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if got.BaseDamage != tt.want.BaseDamage {
					t.Errorf("[%s] BaseDamage = %v, want %v", tt.id, got.BaseDamage, tt.want.BaseDamage)
				}
				if got.DefenseReduced != tt.want.DefenseReduced {
					t.Errorf("[%s] DefenseReduced = %v, want %v", tt.id, got.DefenseReduced, tt.want.DefenseReduced)
				}
				if got.FinalDamage != tt.want.FinalDamage {
					t.Errorf("[%s] FinalDamage = %v, want %v", tt.id, got.FinalDamage, tt.want.FinalDamage)
				}
			}
		})
	}
}

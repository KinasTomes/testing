package combat

import (
	"reflect"
	"testing"
)

func TestCalculateDamage_AllUses(t *testing.T) {
	tests := []struct {
		name        string
		input       CombatInput
		wantResult  CombatResult
		wantErr     bool
		Description string
	}{
		{
			name: "Validation error - Attack too low",
			input: CombatInput{
				Attack:    0,
				Defense:   500,
				AmmoType:  AmmoNormal,
				ArmorType: ArmorMedium,
			},
			wantResult:  CombatResult{},
			wantErr:     true,
			Description: "Cover uses of input and err on the validation error path",
		},
		{
			name: "Successful calculation - Without damage reduction floor limit",
			input: CombatInput{
				Attack:    100,
				Defense:   500,
				AmmoType:  AmmoHE,
				ArmorType: ArmorLight,
			},
			wantResult: CombatResult{
				BaseDamage:     100.0,
				DefenseReduced: 0.5,
				FinalDamage:    60,
			},
			wantErr:     false,
			Description: "Cover uses of all variables along the normal success path",
		},
		{
			name: "Successful calculation - With damage reduction floor limit",
			input: CombatInput{
				Attack:    100,
				Defense:   950,
				AmmoType:  AmmoAP,
				ArmorType: ArmorHeavy,
			},
			wantResult: CombatResult{
				BaseDamage:     100.0,
				DefenseReduced: 0.1, // floor clamped from 0.05
				FinalDamage:    13,
			},
			wantErr:     false,
			Description: "Cover true branch of defRatio < 0.1 floor limit",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := CalculateDamage(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateDamage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("CalculateDamage() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

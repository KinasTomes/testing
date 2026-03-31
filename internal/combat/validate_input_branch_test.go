package combat

import (
	"testing"
)

func TestValidateInput_BranchCoverage(t *testing.T) {
	tests := []struct {
		id      string
		path    string
		input   CombatInput
		wantErr bool
	}{
		{
			id:   "V1",
			path: "1T, 6",
			input: CombatInput{
				Attack:    0,
				Defense:   100,
				AmmoType:  AmmoNormal,
				ArmorType: ArmorLight,
			},
			wantErr: true,
		},
		{
			id:   "V2",
			path: "1F, 2T, 7",
			input: CombatInput{
				Attack:    100,
				Defense:   1001,
				AmmoType:  AmmoNormal,
				ArmorType: ArmorLight,
			},
			wantErr: true,
		},
		{
			id:   "V3",
			path: "1F, 2F, 3F, 8",
			input: CombatInput{
				Attack:    100,
				Defense:   100,
				AmmoType:  "Laser",
				ArmorType: ArmorLight,
			},
			wantErr: true,
		},
		{
			id:   "V4",
			path: "1F, 2F, 3T, 4F, 9",
			input: CombatInput{
				Attack:    100,
				Defense:   100,
				AmmoType:  AmmoAP,
				ArmorType: "Diamond",
			},
			wantErr: true,
		},
		{
			id:   "V5",
			path: "1F, 2F, 3T, 4T, 5",
			input: CombatInput{
				Attack:    100,
				Defense:   100,
				AmmoType:  AmmoHE,
				ArmorType: ArmorLight,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			err := ValidateInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("[%s] ValidateInput() error = %v, wantErr %v", tt.id, err, tt.wantErr)
			}
		})
	}
}

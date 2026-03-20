package combat

import (
	"testing"
)

// TestCalculateDamage_DecisionTable kiểm thử hàm CalculateDamage theo phương pháp Bảng Quyết Định.
//
// Các điều kiện (conditions) được phân tích:
//   C1 - AmmoType  : Normal | HE | AP | invalid
//   C2 - ArmorType : Light | Medium | Heavy | invalid
//   C3 - Attack    : hợp lệ [1..5000] | ngoài biên (<1 hoặc >5000)
//   C4 - Defense   : hợp lệ [0..1000] | ngoài biên (<0 hoặc >1000)
//   C5 - Defense cao (>= 900): kích hoạt sàn MinDamageReductionFloor = 0.1
//
// Kết quả hành động (actions):
//   A1 - Trả về CombatResult hợp lệ
//   A2 - Trả về error (validate thất bại)

type dtCase struct {
	id        string
	name      string
	input     CombatInput
	wantErr   bool
	wantFinal int
	wantBase  float64
	wantDefR  float64
}

func TestCalculateDamage_DecisionTable(t *testing.T) {
	tests := []dtCase{
		// ──────────────────────────────────────────────────────────────────
		// NHÓM 1: Hợp lệ — 9 tổ hợp AmmoType x ArmorType
		// Attack=1000, Defense=200 => defenseRatio = 1 - 200/1000 = 0.8
		// ──────────────────────────────────────────────────────────────────
		{
			id: "DT01", name: "Normal x Light (C1=Normal, C2=Light, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// armorMod=1.0 => floor(1000 * 1.0 * 0.8) = 800
			wantBase: 1000, wantDefR: 0.8, wantFinal: 800,
		},
		{
			id: "DT02", name: "Normal x Medium (C1=Normal, C2=Medium, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoNormal, ArmorType: ArmorMedium},
			// armorMod=0.8 => floor(1000 * 0.8 * 0.8) = 640
			wantBase: 1000, wantDefR: 0.8, wantFinal: 640,
		},
		{
			id: "DT03", name: "Normal x Heavy (C1=Normal, C2=Heavy, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoNormal, ArmorType: ArmorHeavy},
			// armorMod=0.6 => floor(1000 * 0.6 * 0.8) = 480
			wantBase: 1000, wantDefR: 0.8, wantFinal: 480,
		},
		{
			id: "DT04", name: "HE x Light (C1=HE, C2=Light, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoHE, ArmorType: ArmorLight},
			// armorMod=1.2 => floor(1000 * 1.2 * 0.8) = 960
			wantBase: 1000, wantDefR: 0.8, wantFinal: 960,
		},
		{
			id: "DT05", name: "HE x Medium (C1=HE, C2=Medium, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoHE, ArmorType: ArmorMedium},
			// armorMod=1.1 => floor(1000 * 1.1 * 0.8) = 880
			wantBase: 1000, wantDefR: 0.8, wantFinal: 880,
		},
		{
			id: "DT06", name: "HE x Heavy (C1=HE, C2=Heavy, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoHE, ArmorType: ArmorHeavy},
			// armorMod=0.9 => floor(1000 * 0.9 * 0.8) = 720
			wantBase: 1000, wantDefR: 0.8, wantFinal: 720,
		},
		{
			id: "DT07", name: "AP x Light (C1=AP, C2=Light, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoAP, ArmorType: ArmorLight},
			// armorMod=0.8 => floor(1000 * 0.8 * 0.8) = 640
			wantBase: 1000, wantDefR: 0.8, wantFinal: 640,
		},
		{
			id: "DT08", name: "AP x Medium (C1=AP, C2=Medium, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoAP, ArmorType: ArmorMedium},
			// armorMod=1.2 => floor(1000 * 1.2 * 0.8) = 960
			wantBase: 1000, wantDefR: 0.8, wantFinal: 960,
		},
		{
			id: "DT09", name: "AP x Heavy (C1=AP, C2=Heavy, C3=ok, C4=ok)",
			input: CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoAP, ArmorType: ArmorHeavy},
			// armorMod=1.3 => floor(1000 * 1.3 * 0.8) = 1040
			wantBase: 1000, wantDefR: 0.8, wantFinal: 1040,
		},

		// ──────────────────────────────────────────────────────────────────
		// NHÓM 2: Sàn phòng thủ (C5 kích hoạt — Defense >= 900)
		// ──────────────────────────────────────────────────────────────────
		{
			id: "DT10", name: "Defense=950 kich hoat san 0.1 (C5=true, AP x Heavy)",
			input: CombatInput{Attack: 1000, Defense: 950, AmmoType: AmmoAP, ArmorType: ArmorHeavy},
			// defenseRatio = max(1-0.95, 0.1) = max(0.05, 0.1) = 0.1
			// floor(1000 * 1.3 * 0.1) = 130
			wantBase: 1000, wantDefR: 0.1, wantFinal: 130,
		},
		{
			id: "DT11", name: "Defense=1000 san toi da (C5=true, Normal x Light)",
			input: CombatInput{Attack: 1000, Defense: 1000, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			// defenseRatio = max(0.0, 0.1) = 0.1
			// floor(1000 * 1.0 * 0.1) = 100
			wantBase: 1000, wantDefR: 0.1, wantFinal: 100,
		},

		// ──────────────────────────────────────────────────────────────────
		// NHÓM 3: Lỗi — Attack ngoài biên (C3 vi phạm)
		// ──────────────────────────────────────────────────────────────────
		{
			id: "DT12", name: "Attack=0 duoi bien min (C3=invalid)",
			input:   CombatInput{Attack: 0, Defense: 200, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			wantErr: true,
		},
		{
			id: "DT13", name: "Attack=5001 vuot bien max (C3=invalid)",
			input:   CombatInput{Attack: 5001, Defense: 200, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			wantErr: true,
		},

		// ──────────────────────────────────────────────────────────────────
		// NHÓM 4: Lỗi — Defense ngoài biên (C4 vi phạm)
		// ──────────────────────────────────────────────────────────────────
		{
			id: "DT14", name: "Defense=-1 duoi bien min (C4=invalid)",
			input:   CombatInput{Attack: 1000, Defense: -1, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			wantErr: true,
		},
		{
			id: "DT15", name: "Defense=1001 vuot bien max (C4=invalid)",
			input:   CombatInput{Attack: 1000, Defense: 1001, AmmoType: AmmoNormal, ArmorType: ArmorLight},
			wantErr: true,
		},

		// ──────────────────────────────────────────────────────────────────
		// NHÓM 5: Lỗi — AmmoType / ArmorType không hợp lệ (C1/C2 invalid)
		// ──────────────────────────────────────────────────────────────────
		{
			id: "DT16", name: "AmmoType khong hop le = LASER (C1=invalid)",
			input:   CombatInput{Attack: 1000, Defense: 200, AmmoType: "LASER", ArmorType: ArmorLight},
			wantErr: true,
		},
		{
			id: "DT17", name: "ArmorType khong hop le = Mithril (C2=invalid)",
			input:   CombatInput{Attack: 1000, Defense: 200, AmmoType: AmmoNormal, ArmorType: "Mithril"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.id+"_"+tt.name, func(t *testing.T) {
			got, err := CalculateDamage(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("[%s] FAIL: mong doi error nhung khong nhan duoc error", tt.id)
				}
				return
			}

			if err != nil {
				t.Fatalf("[%s] FAIL: khong mong doi error nhung nhan duoc: %v", tt.id, err)
			}
			if got.FinalDamage != tt.wantFinal {
				t.Errorf("[%s] FAIL: FinalDamage want=%d got=%d", tt.id, tt.wantFinal, got.FinalDamage)
			}
			if got.BaseDamage != tt.wantBase {
				t.Errorf("[%s] FAIL: BaseDamage want=%.0f got=%.0f", tt.id, tt.wantBase, got.BaseDamage)
			}
			if got.DefenseReduced != tt.wantDefR {
				t.Errorf("[%s] FAIL: DefenseReduced want=%.4f got=%.4f", tt.id, tt.wantDefR, got.DefenseReduced)
			}
		})
	}
}

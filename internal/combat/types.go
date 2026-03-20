package combat

import "errors"

type ArmorType string

const (
	ArmorLight  ArmorType = "Light"
	ArmorMedium ArmorType = "Medium"
	ArmorHeavy  ArmorType = "Heavy"
)

// AmmoType đại diện cho loại đạn
type AmmoType string

const (
	AmmoAP     AmmoType = "AP"     // Armor Piercing - Hiệu quả vs Heavy Armor
	AmmoHE     AmmoType = "HE"     // High Explosive - Hiệu quả vs Light Armor
	AmmoNormal AmmoType = "Normal" // Đạn thường - cân bằng
)

// Giới hạn hợp lệ của chỉ số.
const (
	MinAttack               = 1
	MaxAttack               = 5000
	MinDefense              = 0
	MaxDefense              = 1000
	MinDamageReductionFloor = 0.1 // Giảm tối thiểu 10% sát thương (tức nhận tối đa 90%)
)

type CombatInput struct {
	Attack    int
	Defense   int
	AmmoType  AmmoType
	ArmorType ArmorType
}

type CombatResult struct {
	BaseDamage     float64 // Sát thương cơ bản trước modifier
	DefenseReduced float64 // Hệ số phòng thủ địch làm giảm damage
	FinalDamage    int     // Sát thương cuối cùng (làm tròn xuống)
}

func ValidateInput(input CombatInput) error {
	if input.Attack < MinAttack || input.Attack > MaxAttack {
		return errors.New("attack phải nằm trong khoảng [1, 5000]")
	}
	if input.Defense < MinDefense || input.Defense > MaxDefense {
		return errors.New("defense phải nằm trong khoảng [0, 1000]")
	}
	if _, ok := ArmorEfficiency[input.AmmoType]; !ok {
		return errors.New("ammoType không hợp lệ, chỉ chấp nhận: AP, HE, Normal")
	}
	if _, ok := ArmorEfficiency[input.AmmoType][input.ArmorType]; !ok {
		return errors.New("armorType không hợp lệ, chỉ chấp nhận: Light, Medium, Heavy")
	}
	return nil
}

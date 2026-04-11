# Báo cáo kiểm thử độ phủ all-uses

## 1. Phân tích các def, use của từng biến.

**Biến `input`**
*   `def(input)` = { 0 }
*   `c-use(input)` = { 1, 3, 5 }
*   `p-use(input)` = { }

**Biến `err`**
*   `def(err)` = { 1 }
*   `c-use(err)` = { 7 }
*   `p-use(err)` = { 2 }

**Biến `armorMod`**
*   `def(armorMod)` = { 3 }
*   `c-use(armorMod)` = { 5 }
*   `p-use(armorMod)` = { }

**Biến `defRatio`**
*   `def(defRatio)` = { 3, 8 }
*   `c-use(defRatio)` = { 5, 6 }
*   `p-use(defRatio)` = { 4 }

**Biến `baseDmg`**
*   `def(baseDmg)` = { 5 }
*   `c-use(baseDmg)` = { 5, 6 }
*   `p-use(baseDmg)` = { }

**Biến `finalDmg`**
*   `def(finalDmg)` = { 5 }
*   `c-use(finalDmg)` = { 6 }
*   `p-use(finalDmg)` = { }

---

## 2. Xác Định Các Du-pairs
*   **`input`**: (0, 1), (0, 3), (0, 5)
*   **`err`**: (1, 2T), (1, 2F), (1, 7), (1, 3)
*   **`armorMod`**: (3, 5)
*   **`defRatio`**: (3, 4T), (3, 4F), (3, 5), (3, 6), (8, 5), (8, 6)
*   **`baseDmg`**: (5, 5), (5, 6)
*   **`finalDmg`**: (5, 6)

---

## 3. Sinh các ca kiểm thử

| Test Case | Input | Complete Path | Các du-pairs được phủ |
| :--- | :--- | :--- | :--- |
| **TC1**<br>*(Validation error)* | `Atk: 0`<br>`Def: 500`<br>`Ammo: Normal`<br>`Armor: Medium` | 0 -> 1 -> 2T -> 7 | **`input`**: `(0, 1)`<br>**`err`**: `(1, 7)`, `(1, (2, 7))` |
| **TC2**<br>*(Success, No floor limit)* | `Atk: 100`<br>`Def: 500`<br>`Ammo: HE`<br>`Armor: Light` | 0 -> 1 -> 2F -> 3 -> 4F -> 5 -> 6 | **`input`**: `(0, 3)`, `(0, 5)`<br>**`err`**: `(1, (2, 3))`<br>**`armorMod`**: `(3, 5)`<br>**`defRatio`**: `(3, 5)`, `(3, 6)`, `(3, (4, 5))`<br>**`baseDmg`**: `(5, 5)`, `(5, 6)`<br>**`finalDmg`**: `(5, 6)` |
| **TC3**<br>*(Success, Hits floor limit)* | `Atk: 100`<br>`Def: 950`<br>`Ammo: AP`<br>`Armor: Heavy` | 0 -> 1 -> 2F -> 3 -> 4T -> 8 -> 5 -> 6 | **`defRatio`**: `(3, (4, 8))`, `(8, 5)`, `(8, 6)` |

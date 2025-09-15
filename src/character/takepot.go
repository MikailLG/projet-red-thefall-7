package character

import (
	"fmt"
)

func removeItem(personnage Character, item string) bool {
	for i, it := range personnage.Inventaire {
		if it == item {
			personnage.Inventaire = append(personnage.Inventaire[:i], personnage.Inventaire[i+1:]...)
			return true
		}
	}
	return false
}

func takePot(personnage Character) {
	if !removeItem(personnage, "Potion de vie") {
		fmt.Println("Pas de potion de vie dans l'inventaire.")
		return
	}
	personnage.HP += 50
	if personnage.HP > p.MaxHP {
		personnage.HP = p.MaxHP
	}
	fmt.Printf("Potion utilisée. PV : %d / %d\n", personnage.HP, personnage.MaxHP)
}

package character

import (
	"fmt"
)

func RemoveItem(personnage Character, item string) bool {
	for i, it := range personnage.Inventaire {
		if it == item {
			personnage.Inventaire = append(personnage.Inventaire[:i], personnage.Inventaire[i+1:]...)
			return true
		}
	}
	return false
}

func TakePot(personnage Character) {
	if !RemoveItem(personnage, "Potion de vie") {
		fmt.Println("Pas de potion de vie dans l'inventaire.")
		return
	}
	personnage.PointDeVie += 50
	if personnage.PointDeVie > personnage.PointDeVieMax {
		personnage.PointDeVie = personnage.PointDeVieMax
	}
	fmt.Printf("Potion utilisée. PV : %d / %d\n", personnage.PointDeVie, personnage.PointDeVieMax)
}

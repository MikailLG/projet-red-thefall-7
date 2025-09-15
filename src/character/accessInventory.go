package character

import (
	"fmt"
)

func AccessInventory(personnage Character) {
	if len(personnage.Inventory) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		fmt.Println("Inventaire :")
		for _, objet := range personnage.Inventory {
			fmt.Println("- " + objet)
		}
	}
}

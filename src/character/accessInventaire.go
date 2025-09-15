package character

import (
	"fmt"
)

func accessInventaire(personnage Character) {
	if len(personnage.Inventaire) == 0 {
		fmt.Println("Votre inventaire est vide.")
	} else {
		fmt.Println("Inventaire :")
		for _, objet := range personnage.Inventaire {
			fmt.Println("- " + objet)
		}
	}
}
